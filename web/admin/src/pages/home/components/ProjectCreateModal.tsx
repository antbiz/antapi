import React from 'react';
import { useRequest } from 'umi';
import { Modal, Form, Input, Space, Button, message } from 'antd';
import { createProject } from '@/services/project';

const ProjectCreateModal: React.FC<{
  visible: boolean;
  onSuccess: () => void;
  onClose: () => void;
}> = ({ visible, onClose, onSuccess }) => {
  const { run, loading } = useRequest(
    async (data: any) => {
      await createProject(data);
      onSuccess();
    },
    {
      manual: true,
      onSuccess: () => message.success('创建项目成功'),
    },
  );

  return (
    <Modal
      centered
      title="创建项目"
      footer={null}
      visible={visible}
      onOk={() => onClose()}
      onCancel={() => onClose()}
    >
      <Form
        name="basic"
        layout="vertical"
        labelCol={{ span: 6 }}
        labelAlign="left"
        onFinish={(v = {}) => {
          run(v);
        }}
      >
        <Form.Item
          label="项目 ID"
          name="name"
          rules={[
            { required: true, message: '请输入项目 ID！' },
            {
              pattern: /^[a-zA-Z0-9]{1,16}$/,
              message: '项目 Id 仅支持字母与数字，不大于 16 个字符',
            },
          ]}
        >
          <Input placeholder="项目 ID，如 website，仅支持字母与数字，不大于 16 个字符" />
        </Form.Item>

        <Form.Item
          label="项目名"
          name="title"
          rules={[{ required: true, message: '请输入项目名！' }]}
        >
          <Input placeholder="项目名，如官网" />
        </Form.Item>

        <Form.Item label="项目介绍" name="description">
          <Input placeholder="项目介绍，如官网内容管理" />
        </Form.Item>

        <Form.Item>
          <Space size="large" style={{ width: '100%', justifyContent: 'flex-end' }}>
            <Button onClick={() => onClose()}>取消</Button>
            <Button type="primary" htmlType="submit" loading={loading}>
              创建
            </Button>
          </Space>
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default ProjectCreateModal;
