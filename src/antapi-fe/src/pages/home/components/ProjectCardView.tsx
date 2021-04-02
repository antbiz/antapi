import React from 'react';
import { history, useAccess } from 'umi';
import { Card, Tooltip, Typography } from 'antd';
import { PlusCircleTwoTone } from '@ant-design/icons';
import './index.less';

const { Title, Paragraph } = Typography;

const cardStyle = {
  height: '80px',
  borderRadius: '10px',
  boxShadow: '0 3px 6px rgb(12 18 100 / 6%)',
};

const cardBodyStyle = {
  padding: '15px',
};

const CreateProject: React.FC<{ onClick: () => void }> = ({ onClick }) => {
  return (
    <div className="project-item" onClick={onClick}>
      <Card bordered={false} style={cardStyle} bodyStyle={cardBodyStyle}>
        <div className="project-icon" style={{ backgroundColor: '#fff' }}>
          <PlusCircleTwoTone style={{ fontSize: '46px' }} />
        </div>
      </Card>
      <div className="ml-5 flex-1">
        <Typography.Title level={4} ellipsis={{ rows: 2, expandable: false }}>
          创建新项目
        </Typography.Title>
      </div>
    </div>
  );
};

/**
 * 项目 card 视图
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
    <div className="project-container">
      {projects.map((project) => (
        <div
          className="project-item"
          key={`${project.id}`}
          onClick={() => {
            history.push(`/project/${project.id}/overview`);
          }}
        >
          <Card bordered={false} style={cardStyle} bodyStyle={cardBodyStyle}>
            <div className="project-icon flex items-center">{project.name.slice(0, 1)}</div>
          </Card>
          <div className="ml-5 flex-1" style={{ maxWidth: '140px' }}>
            <Tooltip title={project.name} placement="topLeft">
              <Title level={4} ellipsis>
                {project.name}
              </Title>
            </Tooltip>
            <Tooltip title={project.description} placement="bottomLeft">
              <Paragraph ellipsis={{ rows: 2, expandable: false }} className="mb-0">
                {project.description || ''}
              </Paragraph>
            </Tooltip>
          </div>
        </div>
      ))}
      {isAdmin && <CreateProject onClick={onCreateProject} />}
    </div>
  );
}
