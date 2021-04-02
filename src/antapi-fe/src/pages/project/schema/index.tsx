import { PlusOutlined } from '@ant-design/icons';
import { Button, Divider, Popconfirm, message, Modal} from 'antd';
import React, { useRef, useState } from 'react';
import { Link } from 'umi';
import { PageContainer } from '@ant-design/pro-layout';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import { getSchemas } from '@/services/schema';
import Generator from 'fr-generator';
import { getProjectId } from '@/utils';

/**
 * 添加
 *
 * @param schema
 */
const handleAdd = async (schema: API.Schema) => {
  const hide = message.loading('正在添加');
  try {
    await addRule({ ...schema });
    hide();
    message.success('添加成功');
    return true;
  } catch (error) {
    hide();
    message.error('添加失败请重试！');
    return false;
  }
};

/**
 * 更新
 *
 * @param fields
 */
// const handleUpdate = async (schema: API.Schema) => {
//   const hide = message.loading('正在更新');
//   try {
//     await updateRule({...schema});
//     hide();
//     message.success('更新成功');
//     return true;
//   } catch (error) {
//     hide();
//     message.error('更新失败请重试！');
//     return false;
//   }
// };

/**
 * 删除
 *
 * @param id
 */
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
  /** 新建窗口的弹窗 */
  const [createModalVisible, handleModalVisible] = useState<boolean>(false);
  /** 更新窗口的弹窗 */
  // const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);

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
        <a
          key="update"
          onClick={() => {
            handleModalVisible(true);
          }}
        >
          更新
        </a>,
        <Popconfirm
          key="delete"
          title="确定删除？"
          okText="是"
          cancelText="否"
          onConfirm={() => {
            handleRemove(record.id);
          }}
        >
          <a href="#">删除</a>
        </Popconfirm>,
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
          <Button
            type="primary"
            key="primary"
            onClick={() => {
              handleModalVisible(true);
            }}
          >
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
      <Modal
        title="新建模型"
        width="90%"
        bodyStyle={{ height: '70vh'}}
        visible={createModalVisible}
        cancelText="取消"
        okText="保存"
        onCancel={() => handleModalVisible(false)}
        onOk={async (value) => {
          const success = await handleAdd(value as API.Schema);
          if (success) {
            handleModalVisible(false);
            if (actionRef.current) {
              actionRef.current.reload();
            }
          }
        }}
      >
        <Generator
          extraButtons={[true, true, false, false]}
        />
      </Modal>
    </PageContainer>
  );
};

export default SchemaList;
