import React, { useEffect, useState } from 'react';
import { getSchemas } from '@/services/schema';
import { getProjectId } from '@/utils';
import { Menu, Spin } from 'antd';
import { useSourceCtx } from '../context';

export default (): React.ReactNode => {
  const { setCurrentSchema } = useSourceCtx();
  const [schemas, setSchemas] = useState<API.Schema[]>([]);
  const projectId = getProjectId();
  const onLoadSchemas = async () => {
    const { data } = await getSchemas({ projectId });
    if (data.length > 0) {
      setCurrentSchema(data[0]);
      setSchemas(data);
    }
  };

  useEffect(() => {
    onLoadSchemas();
  }, []);

  return schemas?.length ? (
    <Menu
      defaultSelectedKeys={[schemas[0].collectionName]}
      mode="inline"
      onClick={(item) => {
        setCurrentSchema(schemas.find((schema) => schema.collectionName === item.key));
      }}
    >
      {schemas.map((schema) => {
        return <Menu.Item key={schema.collectionName}>{schema.displayName}</Menu.Item>;
      })}
    </Menu>
  ) : (
    <Spin tip="加载中..." />
  );
};
