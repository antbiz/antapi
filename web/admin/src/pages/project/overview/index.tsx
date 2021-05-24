import React from 'react';
import { Card, Typography } from 'antd';

export default (): React.ReactNode => {
  return (
    <Card>
      <Typography.Title level={2} style={{ textAlign: 'center' }}>
        欢迎使用 AntApi
      </Typography.Title>
    </Card>
  );
};
