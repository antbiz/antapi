import React, { useEffect, useState } from 'react';
import { getSchemas } from '@/services/schema';
import { getProjectName } from '@/utils';
import { Menu, Spin, Empty } from 'antd';
import { useSourceCtx } from '../context';

const parseSchema = (schema: Record<string, unknown>): Record<string, unknown> => {
  if (schema._properties) {
    schema.properties = JSON.parse(schema._properties);
    delete schema._properties;
  }
  return schema;
};

export default (): React.ReactNode => {
  const { setCurrentSchema } = useSourceCtx();
  const [schemas, setSchemas] = useState<API.Schema[]>(null);
  const projectName = getProjectName();
  const onLoadSchemas = async () => {
    const { data = [] } = await getSchemas({ projectName });
    const schemas = [];
    data.forEach((item) => {
      try {
        schemas.push(parseSchema(item));
      } catch(error) {
        console.error(error);
      }
    });
    if (schemas.length > 0) {
      setCurrentSchema(schemas[0]);
    }
    setSchemas(schemas);
  };

  useEffect(() => {
    onLoadSchemas();
  }, []);

  if (schemas === null) {
    return <Spin tip="加载中..." style={{ marginTop: '30vh', marginLeft: '5vw' }} />;
  }

  return schemas.length ? (
    <Menu
      defaultSelectedKeys={[schemas[0].name]}
      mode="inline"
      onClick={(item) => {
        setCurrentSchema(schemas.find((schema) => schema.name === item.key));
      }}
    >
      {schemas.map((schema) => {
        return <Menu.Item key={schema.name}>{schema.displayName}</Menu.Item>;
      })}
    </Menu>
  ) : (
    <div>
      <Empty
        description="先去创建模型吧"
        image={Empty.PRESENTED_IMAGE_SIMPLE}
        imageStyle={{ height: 60 }}
        style={{ marginTop: '26vh' }}
      />
    </div>
  );
};
