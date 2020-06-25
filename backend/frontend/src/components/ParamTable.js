import React, { forwardRef } from "react";
import MaterialTable from "material-table";
import Paper from "@material-ui/core/Paper";
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
    { title: "ID", field: "ID", type: "numeric", editable: "never" },
    {
      title: "Torque",
      field: "Torque",
      type: "numeric",
      editComponent: (props) => {
        return (
          <TextField
            value={props.value}
            type="number"
            onChange={(e) => props.onChange(e.target.value)}
            inputProps={{ min: "0", max: "65535", step: "1" }}
          />
        );
      },
    },
    {
      title: "AD",
      field: "AD",
      type: "numeric",
      editComponent: (props) => {
        return (
          <TextField
            value={props.value}
            type="number"
            onChange={(e) => props.onChange(e.target.value)}
            inputProps={{ min: "0", max: "65535", step: "1" }}
          />
        );
      },
    },
  ];
  return (
    <Paper>
      <MaterialTable
        options={{
          search: false,
          sorting: false,
          paging: false,
          toolbar: false,
          showFirstLastPageButtons: false,
        }}
        icons={tableIcons}
        localization={{ pagination: { labelRowsPerPage: "12" } }}
        title="Editable Example"
        columns={columns}
        data={props.params}
        editable={{
          isEditable: (rowData) => rowData.name !== "ID",
          isDeleteHidden: (rowData) => true,
          onRowUpdate: props.handleRowUpdate,
        }}
      />
    </Paper>
  );
}
