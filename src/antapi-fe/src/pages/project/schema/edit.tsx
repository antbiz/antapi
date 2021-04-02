import React from 'react';
import Generator from 'fr-generator';
import { PageContainer } from '@ant-design/pro-layout';
import { Card } from 'antd';

const defaultValue = {
  schema: {
    type: 'object',
    properties: {
      inputName: {
        title: '简单输入框',
        type: 'string',
      },
    },
  },
  displayType: 'row',
  showDescIcon: true,
  labelWidth: 120,
};

const SchemaForm: React.FC = () => {
  return (
    <PageContainer
      fixedHeader
    >
      <Card>
        <Generator defaultValue={defaultValue} />
      </Card>
    </PageContainer>
  );
};

export default SchemaForm;
