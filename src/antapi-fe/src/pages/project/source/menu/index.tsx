import React, { useEffect, useState } from 'react';
import { getSchemas } from '@/services/schema';
import { getProjectName } from '@/utils';
import { Menu, Spin, Empty } from 'antd';
import { useSourceCtx } from '../context';

export default (): React.ReactNode => {
  const { setCurrentSchema } = useSourceCtx();
  const [schemas, setSchemas] = useState<API.Schema[]>([]);
  const projectName = getProjectName();
  const onLoadSchemas = async () => {
    const { data } = await getSchemas({ projectName });
    if (data.length > 0) {
      setCurrentSchema(data[0]);
      setSchemas(data);
    }
  };

  useEffect(() => {
    onLoadSchemas();
  }, []);

  if (schemas === null) {
    return (
      <Spin tip="加载中..." style={{ marginTop: '30vh', marginLeft: '5vw' }} />
    )
  }

  return schemas.length ? (
    <Menu
      defaultSelectedKeys={[schemas[0].collectionName]}
      mode="inline"
      onClick={(item) => {
        setCurrentSchema(schemas.find((schema) => schema.collectionName === item.key));
      }}
    >
      {schemas.map((schema) => {
        return <Menu.Item key={schema.collectionName}>{schema.title}</Menu.Item>;
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
