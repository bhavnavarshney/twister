import React from "react";
import NumericInput from "react-numeric-input";

export default function UIntInput(props) {
  const { inputRef, ...other } = props;
  return (
    <NumericInput
      ref={(ref) => {
        inputRef(ref ? ref.inputElement : null);
      }}
      {...other}
      strict
      min={0}
      max={65535}
      step={1}
      style={{
        "input:not(.form-control)": {
          border: "none",
        }
      }}
    />
  );
}
