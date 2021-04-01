import React from 'react';
import { Menu, Row, Col } from 'antd';

export type SchemaMenuListProps = {
  // ref: React.MutableRefObject<API.Schema[] | []>;
  currentSchemaId?: string;
  onSelect?: (key: string) => void;
  schemas: API.Schema[];
}

/**
 * 展示模型列表
 */
const SchemaMenuList: React.FC<SchemaMenuListProps> = ({
  currentSchemaId, onSelect, schemas
}) => {
  // const defaultSelectedMenu = currentSchema?.id ? [currentSchema.id] : [];
  return (
    schemas?.length ? (
      <Menu
        mode="inline"
        defaultSelectedKeys={currentSchemaId || schemas[0].id}
        onClick={({ key }) => {
          onSelect(key);
        }}
      >
        {schemas.map((item: Schema) => (
          <Menu.Item key={item.id}>{item.displayName}</Menu.Item>
        ))}
      </Menu>
    ) : (
      <Row justify="center">
        <Col>模型为空</Col>
      </Row>
    )
  )
};

export default SchemaMenuList;
