import React from 'react';
import { Layout, Row, Col, Button, Empty, Space, Typography } from 'antd';
import SchemaToolbar from './SchemaToolbar';
import SchemaFieldPicker from './SchemaFieldPicker';
import SchemaFieldList from './SchemaFieldList';

const SchemaContent: React.FC<{currentSchema: API.Schema}> = ({currentSchema}) => {

  console.log(currentSchema);

  return (
    <>
      <Layout className="full-height schema-layout-content">
        {currentSchema?.id ? (
          <Row>
            <Col flex="1 1 auto" />
            <Col flex="0 1 600px">
              {/* 工具栏 */}
              <Space className="schema-layout-header">
                <Typography.Title level={3}>{currentSchema.displayName}</Typography.Title>
                <SchemaToolbar />
              </Space>

              {/* 字段列表 */}
              <Layout>
                <SchemaFieldList />
              </Layout>
            </Col>
            <Col flex="1 1 auto" />
          </Row>
        ) : (
          <div className="schema-empty">
            <Empty description="创建你的模型，开始使用 CMS">
              <Button
                type="primary"
                onClick={() => {
                }}
              >
                创建模型
              </Button>
            </Empty>
          </div>
        )}
      </Layout>

      {/* 右侧字段类型列表 */}
      <SchemaFieldPicker />
    </>
  );
};

export default SchemaContent;
