import React from 'react';
import ProCard from '@ant-design/pro-card';
import { Tabs } from 'antd';
import ProjectInfo from './project';
import PermissionInfo from './permission';

export default (): React.ReactNode => {
  return (
    <ProCard>
      <Tabs tabPosition="left" defaultActiveKey="0">
        <Tabs.TabPane tab="项目" key="0">
          <ProjectInfo />
        </Tabs.TabPane>
        <Tabs.TabPane tab="权限" key="1">
          <PermissionInfo />
        </Tabs.TabPane>
      </Tabs>
    </ProCard>
  );
};
