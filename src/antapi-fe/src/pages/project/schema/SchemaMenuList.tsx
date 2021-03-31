import React from 'react';
import { Row, Col } from 'antd';

/**
 * 展示模型列表
 */
const SchemaMenuList: React.FC = () => {

  // const defaultSelectedMenu = currentSchema?.id ? [currentSchema.id] : [];

  return (
    <Row justify="center">
      <Col>模型为空</Col>
    </Row>
  )
};

export default SchemaMenuList;
