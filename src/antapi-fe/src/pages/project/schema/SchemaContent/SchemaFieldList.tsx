import React from 'react';
import { Button, Empty, Space } from 'antd';
import { SchemaFieldRender } from './FieldItemRender';

const SchemaFields: React.FC<{currentSchema: API.Schema}> = ({currentSchema}) => {
  return currentSchema?.fields?.length ? (
    <SchemaFieldRender
      schema={currentSchema}
      onFiledClick={(field) => editFiled(field)}
      actionRender={(field) => (
        <Space>
          <Button
            size="small"
            type="primary"
            onClick={(e) => {
              e.stopPropagation();
              editFiled(field);
            }}
          >
            编辑
          </Button>
          <Button
            danger
            size="small"
            type="primary"
            onClick={(e) => {
              e.stopPropagation();
            }}
          >
            删除
          </Button>
        </Space>
      )}
    />
  ) : (
    <div className="schema-empty">
      <Empty description="点击右侧字段类型，添加一个字段" />
    </div>
  );
};

export default SchemaFields;
