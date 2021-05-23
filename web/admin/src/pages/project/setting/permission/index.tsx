import React, { useState, useEffect } from 'react';
import type { ProColumns } from '@ant-design/pro-table';
import { EditableProTable } from '@ant-design/pro-table';
import { getPermissions, updatePermission } from '@/services/project';
import { Skeleton, Alert, message } from 'antd';
import { getProjectName } from '@/utils';

const permFieldProps = {
  fieldProps: {
    options: [
      {
        label: '仅系统用户',
        value: 0,
      },
      {
        label: '仅创建者',
        value: 1,
      },
      {
        label: '仅登录用户',
        value: 2,
      },
      {
        label: '所有用户',
        value: 3,
      },
    ],
  },
  formItemProps: {
    rules: [
      {
        required: true,
        message: '此项为必填项',
      },
    ],
  },
}

export default (): React.ReactNode => {
  const [permissions, setPermissions] = useState<API.Permission[]>(null);
  const columns: ProColumns<API.Permission>[] = [
    {
      title: '模型名称',
      dataIndex: 'title',
      width: '16%',
      editable: false,
    },
    {
      title: '创建权限',
      key: 'createLevel',
      dataIndex: 'createLevel',
      valueType: 'select',
      ...permFieldProps,
    },
    {
      title: '读取权限',
      key: 'readLevel',
      dataIndex: 'readLevel',
      valueType: 'select',
      ...permFieldProps,
    },
    {
      title: '更新权限',
      key: 'updateLevel',
      dataIndex: 'updateLevel',
      valueType: 'select',
      ...permFieldProps,
    },
    {
      title: '删除权限',
      key: 'deleteLevel',
      dataIndex: 'deleteLevel',
      valueType: 'select',
      ...permFieldProps,
    },
    {
      title: '操作',
      valueType: 'option',
      width: 150,
      render: (text, record, _, action) => [
        <a
          key="editable"
          onClick={() => {
            action?.startEditable?.(record._id);
          }}
        >
          编辑
        </a>,
      ],
    },
  ];

  const projectName = getProjectName();
  const onLoadPermissions = async () => {
    const { data = [] } = await getPermissions({ projectName });
    setPermissions(data);
  };
  const savePermission = async (id: string, data: API.Permission) => {
    const hide = message.loading('正在更新');
    try {
      await updatePermission(id, data);
      hide();
      message.success('更新成功');
      return true;
    } catch (error) {
      hide();
    }
  }

  useEffect(() => {
    onLoadPermissions();
  }, []);

  if (permissions === null) {
    return <Skeleton />;
  }

  return (
    <EditableProTable<API.Permission>
      rowKey="_id"
      headerTitle="API 访问权限"
      maxLength={20}
      recordCreatorProps={false}
      type="single"
      columns={columns}
      value={permissions}
      editable={{
        actionRender: (row, config, dom) => [
          <a
            key="save"
            onClick={async () => {
              const data = (await config?.form?.validateFields()) as API.Permission;
              const success = await savePermission(row._id, data[row._id]);
              if (success) {
                const perms = [];
                permissions.forEach((v) => {
                  if (v._id === row._id) {
                    v = {...v, ...data[row._id]}
                  }
                  perms.push(v);
                })
                setPermissions(perms)
                await config?.onSave?.(config.recordKey);
              }
            }}
          >
            保存
          </a>,
          dom.cancel,
        ],
      }}
      toolBarRender={() => {
        return [
          <Alert
            message="接口访问权限仅对非系统用户而言，系统用户对所有接口拥有全部权限！"
            type="info"
            showIcon
            closable
          />
        ];
      }}
    />
  );
};
