import React, { useState, useEffect } from 'react';
import { Form, message } from 'antd';
import ProForm, { ProFormText } from '@ant-design/pro-form';
import { LockOutlined } from '@ant-design/icons';
import { updatePassword } from '@/services/account';

export default (): React.ReactNode => {
  return (
    <ProForm
      onFinish={async (values) => {
        const hide = message.loading('正在更新');
        try {
          await updatePassword(values);
          hide();
          message.success('密码已修改');
        } catch (error) {
          hide();
        }
      }}
      submitter={{
        searchConfig: {
          submitText: '修改密码',
        },
        render: (_, dom) => dom.pop(),
      }}
    >
      <ProFormText.Password
        name="oldPassword"
        label="旧密码"
        width="md"
        placeholder="请输入旧密码"
        fieldProps={{
          prefix: <LockOutlined />,
        }}
        rules={[
          {
            required: true,
            message: '请输入旧密码!',
          },
        ]}
      />
      <ProFormText.Password
        name="password"
        label="新密码"
        width="md"
        placeholder="请输入新密码"
        fieldProps={{
          prefix: <LockOutlined />,
        }}
        rules={[
          {
            required: true,
            message: '请输入新密码!',
          },
        ]}
      />
      <Form.Item noStyle shouldUpdate>
        {(form) => {
          return (
            <ProFormText.Password
              name="password2"
              label="确认密码"
              width="md"
              placeholder="请二次输入新密码"
              fieldProps={{
                prefix: <LockOutlined />,
              }}
              rules={[
                {
                  required: true,
                  message: '请二次输入新密码!',
                },
                {
                  validator: (rule, value, callback) => {
                    if (form.getFieldValue('password') !== value) {
                      callback(new Error('两次密码输入不一致'));
                    } else {
                      callback();
                    }
                  },
                },
              ]}
            />
          );
        }}
      </Form.Item>
    </ProForm>
  );
};
