import React, { useState } from 'react';
import { GridContent } from "@ant-design/pro-layout";
import { Menu } from "antd";
import BaseView from './base';
import PasswordView from './password';
import styles from './index.less';

export default (): React.ReactNode => {
  const [selectKey, setSelectKey] = useState('base');
  const menuMap = {
    base: '基本设置',
    password: '修改密码',
  }

  const renderChildren = () => {
    switch (selectKey) {
      case 'base':
        return <BaseView />;
      case 'password':
        return <PasswordView />;
      default:
        break;
    }
    return null;
  };

  const getMenu = () =>
    Object.keys(menuMap).map((item) => <Menu.Item key={item}>{menuMap[item]}</Menu.Item>);

  return (
    <GridContent>
      <div className={styles.main}>
        <div className={styles.leftMenu}>
          <Menu
            mode={'inline'}
            selectedKeys={[selectKey]}
            onClick={({ key }) => {
              setSelectKey(key);
            }}
          >
            {getMenu()}
          </Menu>
        </div>
        <div className={styles.right}>
          <div className={styles.title}>{menuMap[selectKey]}</div>
          {renderChildren()}
        </div>
      </div>
    </GridContent>
  );
};
