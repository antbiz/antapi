import React from 'react';
import { useSetState } from 'react-use';
import { useRequest, useAccess } from 'umi';
import { useLocalStorageState } from '@umijs/hooks';
import { setTwoToneColor, AppstoreOutlined, UnorderedListOutlined } from '@ant-design/icons';
import { Tooltip, Typography, Empty } from 'antd';
import ProLayout from '@ant-design/pro-layout';
import { getProjects } from '@/services/project';
import ProjectListView from './components/ProjectListView';
import ProjectCardView from './components/ProjectCardView';
import ProjectCreateModal from './components/ProjectCreateModal';
import RightContent from '@/components/RightContent';
import defaultSettings from '../../../config/defaultSettings';
import './index.less';

// 设置图标颜色
setTwoToneColor('#0052d9');

const toggleIconStyle: React.CSSProperties = {
  fontSize: '1.6em',
  fontWeight: 'bold',
  color: '#0052d9',
};

export default (): React.ReactNode => {
  // 布局设置持久化到本地
  const [currentLayout, setLocalLayout] = useLocalStorageState('ANTAPI_CMS_PROJECT_LAYOUT', 'card');
  const [{ modalVisible, reload }, setState] = useSetState({
    reload: 0,
    modalVisible: false,
  });
  const { isAdmin } = useAccess();

  // 请求数据
  const { data = [] } = useRequest(() => getProjects(), {
    refreshDeps: [reload],
  });

  const createProject = () =>
    setState({
      modalVisible: true,
    });

  return (
    <ProLayout
      {...defaultSettings}
      disableContentMargin
      rightContentRender={() => {
        return <RightContent />;
      }}
      menuRender={false}
    >
      <div className="home">
        <div className="flex items-center justify-between mb-10">
          <Typography.Title level={3}>我的项目</Typography.Title>
          <Tooltip title="切换布局">
            <div
              className="toggle-icon flex items-center justify-between cursor-pointer"
              onClick={() => {
                setLocalLayout(currentLayout === 'card' ? 'list' : 'card');
              }}
            >
              {currentLayout === 'card' ? (
                <UnorderedListOutlined style={toggleIconStyle} />
              ) : (
                <AppstoreOutlined style={toggleIconStyle} />
              )}
            </div>
          </Tooltip>
        </div>

        {!isAdmin && !data?.length && (
          <Empty description="项目为空，请联系您的管理员为您分配项目！" />
        )}

        {currentLayout === 'card' ? (
          <ProjectCardView projects={data} onCreateProject={createProject} />
        ) : (
          <ProjectListView projects={data} onCreateProject={createProject} />
        )}

        {/* 新项目创建 */}
        {isAdmin && (
          <ProjectCreateModal
            visible={modalVisible}
            onClose={() => setState({ modalVisible: false })}
            onSuccess={() => {
              setState({
                reload: reload + 1,
                modalVisible: false,
              });
            }}
          />
        )}
      </div>
    </ProLayout>
  );
};
