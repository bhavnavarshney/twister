import React, { forwardRef } from "react";
import MaterialTable from "material-table";
import TextField from "@material-ui/core/TextField";
import Check from "@material-ui/icons/Check";
import Clear from "@material-ui/icons/Clear";
import Edit from "@material-ui/icons/Edit";

const tableIcons = {
  Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
  Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
  Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
};

export default function ParamTable(props) {
  const columns = [
    {
      title: "ID",
      field: "ID",
      type: "numeric",
      editable: "never",
      cellStyle: { textAlign: "left", fontWeight: "400" },
      headerStyle: { textAlign: "left" },
    },
    {
      title: "Torque",
      field: "Torque",
      type: "numeric",
      cellStyle: { textAlign: "left" },
      headerStyle: { textAlign: "left" },
    },
    {
      title: "AD",
      field: "AD",
      type: "numeric",
      cellStyle: { textAlign: "left" },
      headerStyle: { textAlign: "left" },
    },
  ];
  return (
    <MaterialTable
      components={{
        EditField: (props) => (
          <TextField
            style={{ float: "right", maxWidth: "120px" }}
            type="number"
            value={props.value === undefined ? "" : props.value}
            onChange={(event) => props.onChange(event.target.value)}
          />
        ),
      }}
      options={{
        search: false,
        sorting: false,
        paging: false,
        toolbar: true,
        showFirstLastPageButtons: false,
        actionsColumnIndex: -1,
      }}
      icons={tableIcons}
      localization={{ pagination: { labelRowsPerPage: "12" } }}
      title={props.title}
      columns={columns}
      data={props.params}
      editable={{
        isEditable: (rowData) => rowData.name !== "ID",
        isDeleteHidden: (rowData) => true,
        onRowUpdate: props.handleRowUpdate,
      }}
    />
  );
}
