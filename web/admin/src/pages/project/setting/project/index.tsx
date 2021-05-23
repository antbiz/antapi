import React, { useState, useEffect } from 'react';
import { getProject, updateProject, deleteProject } from '@/services/project';
import { Divider, Button, Space, Typography, Form, Input, Skeleton, Modal, message } from 'antd';
import { getProjectName } from '@/utils';
import ProCard from '@ant-design/pro-card';

/**
 * 更新模型
 *
 * @param project
 */
const handleUpdate = async (project: API.Project) => {
  const hide = message.loading('正在更新');
  try {
    await updateProject(project.id, project);
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
    await deleteProject(id);
    hide();
    message.success('删除成功，即将刷新');
    return true;
  } catch (error) {
    hide();
    message.error('删除失败，请重试');
    return false;
  }
};

const ProjectDangerAction: React.FC<{ project: API.Project }> = ({ project }) => {
  const [modalVisible, setModalVisible] = useState(false);
  const [projectTitle, setProjectTitle] = useState('');

  return (
    <>
      <Typography.Title level={5}>危险操作</Typography.Title>
      <Divider />
      <Button
        danger
        type="primary"
        onClick={() => {
          setModalVisible(true);
        }}
      >
        删除项目
      </Button>
      <Modal
        centered
        title="删除项目"
        visible={modalVisible}
        onCancel={() => setModalVisible(false)}
        onOk={() => handleDelete(project.id)}
        okButtonProps={{
          disabled: projectTitle !== project.title,
        }}
      >
        <Space direction="vertical">
          <Typography.Paragraph strong>删除项目会删除项目中的内容模型等数据</Typography.Paragraph>
          <Typography.Paragraph strong>
            删除项目是不能恢复的，您确定要删除此项目吗？
            如果您想继续，请在下面的方框中输入此项目的名称：
            <Typography.Text strong mark>
              {project.title}
            </Typography.Text>
          </Typography.Paragraph>
          <Input value={projectTitle} onChange={(e) => setProjectTitle(e.target.value)} />
        </Space>
      </Modal>
    </>
  );
};

export default (): React.ReactNode => {
  const projectName = getProjectName();
  const [changed, setChanged] = useState(false);
  const [project, setProject] = useState<API.Project>();
  const onLoadProject = async () => {
    const { data } = await getProject(projectName);
    if (data) {
      setProject(data);
    }
  };

  useEffect(() => {
    onLoadProject();
  }, []);

  if (!project) {
    return <Skeleton />;
  }

  return (
    <ProCard>
      <Typography.Title level={5}>项目信息</Typography.Title>
      <Divider />
      <Form
        name="basic"
        layout="vertical"
        labelAlign="left"
        initialValues={project}
        onFinish={(v = {}) => {
          handleUpdate(v);
        }}
        onValuesChange={(_, v: Partial<API.Project>) => {
          if (v.name !== project?.name || v.description !== project.description) {
            setChanged(true);
          } else {
            setChanged(false);
          }
        }}
      >
        <Form.Item label="项目 ID">
          <Typography.Paragraph copyable>{project.name}</Typography.Paragraph>
        </Form.Item>
        <Form.Item
          label="项目名称"
          name="title"
          rules={[{ required: true, message: '请输入项目名称！' }]}
        >
          <Input placeholder="项目名，如官网" />
        </Form.Item>

        <Form.Item label="项目介绍" name="description">
          <Input placeholder="项目介绍，如官网内容管理" />
        </Form.Item>

        <Form.Item>
          <Space size="large" style={{ width: '100%', justifyContent: 'flex-end' }}>
            <Button type="primary" htmlType="submit" disabled={!changed}>
              保存
            </Button>
          </Space>
        </Form.Item>
      </Form>

      <ProjectDangerAction project={project} />
    </ProCard>
  );
};
