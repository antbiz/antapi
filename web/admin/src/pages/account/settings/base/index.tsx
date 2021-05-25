import React, { useState, useEffect } from 'react';
import { message } from 'antd';
import { useModel } from 'umi';
import ProForm, { ProFormSelect, ProFormUploadButton } from '@ant-design/pro-form';
import { updateInfo } from '@/services/account';

export default (): React.ReactNode => {
  const { initialState, setInitialState } = useModel('@@initialState');
  const { currentUser } = initialState;
  // FIXME: 用户头像修改
  currentUser.avatar = null;

  return (
    <ProForm
      initialValues={currentUser}
      onFinish={async (values) => {
        const hide = message.loading('正在更新');
        try {
          await updateInfo(values);
          hide();
          message.success('更新成功');
          setInitialState({
            ...initialState,
            currentUser: {
              ...currentUser,
              ...values,
            },
          });
        } catch (error) {
          hide();
        }
      }}
      submitter={{
        searchConfig: {
          submitText: '更新基本信息',
        },
        render: (_, dom) => dom.pop(),
      }}
    >
      <ProFormUploadButton
        name="avatar"
        label="头像"
        max={1}
        fieldProps={{
          name: 'file',
          listType: 'picture-card',
        }}
        action="/upload.do"
      />
      <ProFormSelect
        width="md"
        name="language"
        label="语言"
        rules={[
          {
            required: true,
            message: '请选择语言!',
          },
        ]}
        options={[
          { label: '中文', value: 'zh-CN' },
          { label: 'EN', value: 'en' },
        ]}
      />
    </ProForm>
  );
};
