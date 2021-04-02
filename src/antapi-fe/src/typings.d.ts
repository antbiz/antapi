declare module 'slash2';
declare module '*.css';
declare module '*.less';
declare module '*.scss';
declare module '*.sass';
declare module '*.svg';
declare module '*.png';
declare module '*.jpg';
declare module '*.jpeg';
declare module '*.gif';
declare module '*.bmp';
declare module '*.tiff';
declare module 'omit.js';
declare module 'numeral';
declare module '@antv/data-set';
declare module 'mockjs';
declare module 'react-fittext';
declare module 'bizcharts-plugin-slider';

declare module 'form-render/lib/antd' {
  import React from 'react';

  export interface FRProps {
    schema: object;
    formData: object;
    onChange?: void;
    onMount?: void;
    name?: string;
    column?: number;
    uiSchema?: object;
    widgets?: any;
    FieldUI?: any;
    fields?: any;
    mapping?: object;
    showDescIcon?: boolean;
    showValidate?: boolean;
    displayType?: string;
    onValidate: any;
    readOnly?: boolean;
    labelWidth?: number | string;
  }
  class FormRender extends React.Component<FRProps> {}
  export default FormRender;
}

// google analytics interface
type GAFieldsObject = {
  eventCategory: string;
  eventAction: string;
  eventLabel?: string;
  eventValue?: number;
  nonInteraction?: boolean;
};

type Window = {
  ga: (
    command: 'send',
    hitType: 'event' | 'pageview',
    fieldsObject: GAFieldsObject | string,
  ) => void;
  reloadAuthorized: () => void;
  routerBase: string;
};

declare let ga: () => void;

declare const REACT_APP_ENV: 'test' | 'dev' | 'pre' | false;
