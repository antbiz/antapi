import { PlusOutlined } from '@ant-design/icons';
import { Button, Popconfirm, message, Modal} from 'antd';
import React, { useRef, useState } from 'react';
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
const handleUpdate = async (schema: API.Schema) => {
  const hide = message.loading('正在更新');
  try {
    await updateRule({...schema});
    hide();
    message.success('更新成功');
    return true;
  } catch (error) {
    hide();
    message.error('更新失败请重试！');
    return false;
  }
};

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

export default (): React.ReactNode => {
  /** 新建窗口的弹窗 */
  const [editModalVisible, handleEditModalVisible] = useState<boolean>(false);
  const [currentItem, setCurrentItem] = useState<API.Schema>({});

  const actionRef = useRef<ActionType>();
  const genRef = useRef();
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
        <Popconfirm
          key="delete"
          title="确定删除？"
          okText="是"
          cancelText="否"
          onConfirm={() => {
            handleRemove(record.id);
          }}
        >
          <a style={{color: '#ff7875'}} href="#">删除</a>
        </Popconfirm>,
        <a
          key="update"
          onClick={() => {
            setCurrentItem(record);
            handleEditModalVisible(true);
          }}
        >
          更新
        </a>,
      ],
    },
  ];

  return (
    <>
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
              handleEditModalVisible(true);
            }}
          >
            <PlusOutlined />
            新建
          </Button>,
        ]}
        params={{ projectId }}
        request={getSchemas}
        columns={columns}
      />
      <Modal
        title={currentItem?.id ? "更新模型" : "新建模型"}
        width="90%"
        bodyStyle={{ height: '70vh'}}
        maskClosable={false}
        visible={editModalVisible}
        cancelText="取消"
        okText="保存"
        onCancel={() => handleEditModalVisible(false)}
        onOk={async () => {
          const value = genRef.current.getValue();
          console.log(value)
          const success = value?.id ? await handleUpdate(value as API.Schema) : await handleAdd(value as API.Schema);
          if (success) {
            handleEditModalVisible(false);
            if (actionRef.current) {
              actionRef.current.reload();
            }
          }
        }}
      >
        <Generator
          ref={genRef}
          extraButtons={[true, true, false, false]}
          // defaultValue={currentItem}
        />
      </Modal>
    </>
  );
};
