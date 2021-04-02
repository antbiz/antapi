import { PlusOutlined } from '@ant-design/icons';
import { Button, Divider, Popconfirm, message } from 'antd';
import React, { useRef } from 'react';
import { Link } from 'umi';
import { PageContainer } from '@ant-design/pro-layout';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import { getSchemas } from '@/services/schema';
import { getProjectId } from '@/utils';

// /**
//  * 删除节点
//  *
//  * @param id
//  */
const handleRemove = async (id: string) => {
  const hide = message.loading('正在删除');
  if (!id) return true;
  try {
    await deleteSchema({ id });
    hide();
    message.success('删除成功，即将刷新');
    return true;
  } catch (error) {
    hide();
    message.error('删除失败，请重试');
    return false;
  }
};

const SchemaList: React.FC = () => {
  const actionRef = useRef<ActionType>();
  const projectId = getProjectId();

  const columns: ProColumns<API.Schema>[] = [
    {
      title: '标题',
      dataIndex: 'displayName',
    },
    {
      title: '名称',
      dataIndex: 'collectionName',
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (_, record) => [
        <div key="option">
          <Link to={`/project/${projectId}/schema/edit?id=${record.id}`}>编辑</Link>
          <Divider type="vertical" />
          <Popconfirm
            title="确定删除？"
            okText="是"
            cancelText="否"
            onConfirm={() => {
              handleRemove(record.id);
            }}
          >
            <a href="#">删除</a>
          </Popconfirm>
        </div>,
      ],
    },
  ];

  return (
    <PageContainer>
      <ProTable<API.Schema, API.PageParams>
        actionRef={actionRef}
        rowKey="id"
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [
          <Button type="primary" key="primary" onClick={() => {}}>
            <PlusOutlined />
            新建
          </Button>,
        ]}
        params={{ projectId }}
        request={getSchemas}
        columns={columns}
        rowSelection={{
          onChange: (_, selectedRows) => {
            setSelectedRows(selectedRows);
          },
        }}
      />
    </PageContainer>
  );
};

export default SchemaList;
