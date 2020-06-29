import React from "react";
import InputNumber from 'rc-input-number';

export default function UIntInput(props) {
  const { inputRef, ...other } = props;
  return (
    <InputNumber
      ref={(ref) => {
        inputRef(ref ? ref.inputElement : null);
      }}
      {...other}
      min={0}
      max={65535}
      step={1}
      style={{
        maxWidth: "55px"
      }}
    />
  );
}
