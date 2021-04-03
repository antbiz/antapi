import React, { useState, useRef } from 'react';
import ProCard from '@ant-design/pro-card';
import { Button, message, Modal } from 'antd';
import SourceMenu from './menu';
import Content from './content';
import { SourceCtx } from './context';
import { updateSchema, createSchema, deleteSchema } from '@/services/schema';
import Generator from 'fr-generator';
import './index.less';

/**
 * 添加模型
 *
 * @param schema
 */
const handleAdd = async (schema: API.Schema) => {
  const hide = message.loading('正在添加');
  try {
    await createSchema({ ...schema });
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
 * 更新模型
 *
 * @param schema
 */
const handleUpdate = async (schema: API.Schema) => {
  const hide = message.loading('正在更新');
  try {
    await updateSchema({ ...schema });
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
 * 删除模型
 *
 * @param id
 */
const handleDelete = async (id: string) => {
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
  /** 编辑模型窗口的弹窗 */
  const [editModalVisible, handleEditModalVisible] = useState<boolean>(false);
  const [currentSchema, setCurrentSchema] = useState<API.Schema>();
  const genRef = useRef();

  return (
    <SourceCtx.Provider value={{ currentSchema, setCurrentSchema }}>
      <ProCard split="vertical">
        <ProCard
          colSpan={{
            xs: '40px',
            sm: '80px',
            md: '120px',
            lg: '160px',
            xl: '200px',
          }}
          className="ghost"
        >
          <SourceMenu />
        </ProCard>
        <ProCard
          bordered
          headerBordered
          direction="column"
          gutter={[0, 16]}
          className="content-container"
          extra={
            <>
              <Button
                key="deleteSchema"
                danger
                onClick={() => {
                  handleDelete(currentSchema.id);
                }}
              >
                删除模型
              </Button>
              <Button
                key="updateSchema"
                onClick={() => {
                  handleEditModalVisible(true);
                }}
              >
                更新模型
              </Button>
              <Button
                key="addSchema"
                type="primary"
                onClick={() => {
                  handleEditModalVisible(true);
                }}
              >
                新建模型
              </Button>
            </>
          }
        >
          <Content />
        </ProCard>
      </ProCard>
      <Modal
        title={currentSchema?.id ? '更新模型' : '新建模型'}
        width="90%"
        bodyStyle={{ height: '70vh' }}
        maskClosable={false}
        visible={editModalVisible}
        cancelText="取消"
        okText="保存"
        onCancel={() => handleEditModalVisible(false)}
        onOk={async () => {
          const value = genRef.current.getValue();
          console.log(value);
          const success = value?.id
            ? await handleUpdate(value as API.Schema)
            : await handleAdd(value as API.Schema);
          if (success) {
            handleEditModalVisible(false);
            if (actionRef.current) {
              actionRef.current.reload();
            }
          }
        }}
      >
        <Generator ref={genRef} extraButtons={[true, true, false, false]} />
      </Modal>
    </SourceCtx.Provider>
  );
};
