import { PlusOutlined } from '@ant-design/icons';
import { Button, Popconfirm, message, Modal } from 'antd';
import React, { useRef, useState } from 'react';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import FormRender, { useForm } from 'form-render';
import { createOne, updateOne, deleteOne, getMany } from '@/services/crud';
import { frSchema2ProTableCols } from './utils';
import { useSourceCtx } from '../context';

/**
 * 添加
 *
 * @param collectionName
 * @param data
 */
const handleAdd = async (collectionName: string, data: Record<string, unknown>) => {
  const hide = message.loading('正在添加');
  try {
    await createOne(collectionName, data);
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
 * @param collectionName
 * @param data
 */
const handleUpdate = async (collectionName: string, id: string, data: Record<string, unknown>) => {
  const hide = message.loading('正在更新');
  try {
    await updateOne(collectionName, id, data);
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
 * @param collectionName
 * @param id
 */
const handleDelete = async (collectionName: string, id: string) => {
  const hide = message.loading('正在删除');
  if (!id) return true;
  try {
    await deleteOne(collectionName, id);
    hide();
    message.success('删除成功，即将刷新');
    return true;
  } catch (error) {
    hide();
  }
};

export default (): React.ReactNode => {
  /** 新建窗口的弹窗 */
  const [editModalVisible, handleEditModalVisible] = useState<boolean>(false);
  const [currentItem, setCurrentItem] = useState<API.Schema>({});
  const actionRef = useRef<ActionType>();
  const { currentSchema } = useSourceCtx();
  const form = useForm();

  let columns: ProColumns<Record<string, unknown>>[] = frSchema2ProTableCols(currentSchema);
  columns = columns.concat([
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
            handleDelete(currentSchema.name, record._id);
          }}
        >
          <a style={{ color: '#ff7875' }} href="#">
            删除
          </a>
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
  ]);

  return (
    <>
      <ProTable<Record<string, unknown>, API.PageParams>
        actionRef={actionRef}
        rowKey="_id"
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
        params={{ collectionName: currentSchema.name }}
        request={getMany}
        columns={columns}
      />
      <Modal
        title={
          currentItem?._id ? `更新${currentSchema.displayName}` : `新建${currentSchema.displayName}`
        }
        width="90%"
        bodyStyle={{ height: '70vh' }}
        maskClosable={false}
        visible={editModalVisible}
        cancelText="关闭"
        okText="保存"
        onCancel={() => handleEditModalVisible(false)}
        onOk={form.submit}
      >
        <FormRender
          form={form}
          schema={currentSchema}
          formData={currentItem}
          onFinish={async (data, errors) => {
            if (errors.length == 0) {
              const success = currentItem?._id
                ? await handleUpdate(currentSchema.name, currentItem._id, data)
                : await handleAdd(currentSchema.name, data);
              if (success) {
                handleEditModalVisible(false);
                if (actionRef.current) {
                  actionRef.current.reload();
                }
              }
            }
          }}
        />
      </Modal>
    </>
  );
};
