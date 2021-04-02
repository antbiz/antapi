import React from 'react';
import { history, useAccess } from 'umi';
import ProCard from '@ant-design/pro-card';
import { Tooltip, Typography } from 'antd';
import { PlusCircleOutlined } from '@ant-design/icons';
import './index.less';

const { Title, Paragraph } = Typography;

/**
 * 项目 list 视图
 */
export default function ProjectListView({
  projects,
  onCreateProject,
}: {
  projects: Project[];
  onCreateProject: () => void;
}) {
  const { isAdmin } = useAccess();

  return (
    <ProCard style={{ borderRadius: '5px' }} bodyStyle={{ padding: 0 }}>
      {projects.map((project) => (
        <div
          key={`${project.id}`}
          className="project-list-item flex items-center py-5 px-5"
          onClick={() => {
            history.push(`/project/${project.id}/overview`);
          }}
        >
          <div className="w-2/4 flex items-center">
            <div className="item-icon">{project.name.slice(0, 1)}</div>
            <Tooltip title={project.name}>
              <Title level={5} ellipsis className="ml-5 mb-0" style={{ maxWidth: '80%' }}>
                {project.name}
              </Title>
            </Tooltip>
          </div>
          <Tooltip title={project.description}>
            <Paragraph ellipsis style={{ maxWidth: '80%', marginBottom: 0 }}>
              {project.description || '-'}
            </Paragraph>
          </Tooltip>
        </div>
      ))}
      {isAdmin && (
        <div
          className="project-list-item flex items-center py-5 px-5"
          onClick={() => onCreateProject()}
        >
          <div className="w-2/4 flex items-center">
            <div className="item-icon create-icon">
              <PlusCircleOutlined style={{ fontSize: '24px', color: '#0052d9' }} />
            </div>
            <Title level={5} className="ml-5 mb-0">
              创建新项目
            </Title>
          </div>
        </div>
      )}
    </ProCard>
  );
}
