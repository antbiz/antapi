import React from 'react';
import { Card, Layout, List, message, Typography } from 'antd';
import { FieldTypes } from '@/common';

const SchemaFieldPicker: React.FC = () => {
  return (
    <Layout className="schema-sider" width="220">
      <Typography.Title level={3} className="schema-sider-header">
        内容类型
      </Typography.Title>
      <List
        bordered={false}
        dataSource={FieldTypes}
        renderItem={(item) => (
          <Card
            hoverable
            className="field-card"
            onClick={() => {
              if (!currentSchema) {
                message.info('请选择需要编辑的模型');
              }
            }}
          >
            <List.Item className="item">
              <span>{item.icon}</span>
              <span>{item.name}</span>
            </List.Item>
          </Card>
        )}
      />
    </Layout>
  );
};

export default SchemaFieldPicker;
