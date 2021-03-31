import React, { useEffect } from 'react';
import { history, Link, matchPath } from 'umi';
import RightContent from '@/components/RightContent';
import type { MenuDataItem, BasicLayoutProps } from '@ant-design/pro-layout';
import ProLayout from '@ant-design/pro-layout';
import {
  EyeTwoTone,
  GoldTwoTone,
  DatabaseTwoTone,
  SettingTwoTone,
  setTwoToneColor,
} from '@ant-design/icons';
import defaultSettings from '../../config/defaultSettings';

// 设置图标颜色
setTwoToneColor('#0052d9');

const customMenuData: MenuDataItem[] = [
  {
    authority: 'isAdmin',
    path: '/project/:projectId/overview',
    name: '概览',
    icon: <EyeTwoTone />,
  },
  {
    authority: 'isAdmin',
    path: '/project/:projectId/schema',
    name: '内容模型',
    icon: <GoldTwoTone />,
  },
  {
    authority: 'isAdmin',
    path: '/project/:projectId/content',
    name: '内容集合',
    icon: <DatabaseTwoTone />,
    children: [],
  },
  {
    authority: 'isAdmin',
    path: '/project/:projectId/setting',
    name: '项目设置',
    icon: <SettingTwoTone />,
  },
];

const layoutProps: BasicLayoutProps = {
  // disableContentMargin: true,
  rightContentRender: () => <RightContent />,
  // 面包屑渲染
  itemRender: () => null,
  ...defaultSettings,
};

const ProjectLayout: React.FC<any> = (props) => {
  const { children, location } = props;

  useEffect(() => {
    // 匹配 Path，获取 projectId
    const match = matchPath<{ projectId?: string }>(history.location.pathname, {
      path: '/project/:projectId/*',
      exact: true,
      strict: false,
    });

    // projectId 无效时，重定向到首页
    const { projectId = '' } = match?.params || {};
    if (projectId === ':projectId' || !projectId) {
      history.push('/home');
    }
  }, []);

  return (
    <ProLayout
      menu={{ defaultOpenAll: true }}
      location={location}
      onMenuHeaderClick={() => {
        history.push('/home');
      }}
      menuDataRender={() => {
        return customMenuData;
      }}
      menuItemRender={(menuItemProps, defaultDom) => {
        const match = matchPath<{ projectId?: string }>(history.location.pathname, {
          path: '/project/:projectId/*',
          exact: true,
        });

        // 项目 Id
        const { projectId = '' } = match?.params || {};

        if (menuItemProps.isUrl || menuItemProps.children) {
          return defaultDom;
        }

        if (menuItemProps.path) {
          return (
            <Link
              to={menuItemProps.path.replace(':projectId', projectId)}
            >
              {defaultDom}
            </Link>
          );
        }

        return defaultDom;
      }}
      {...layoutProps}
    >
      {children}
    </ProLayout>
  )
};

export default ProjectLayout;
