import React from 'react';
import { Card, Typography } from 'antd';

export default (): React.ReactNode => {
  return (
    <Card>
      <Typography.Title level={2} style={{ textAlign: 'center' }}>
        这是schema页面
      </Typography.Title>
    </Card>
  );
};
