import React, { forwardRef } from "react";
import MaterialTable from "material-table";
import TextField from "@material-ui/core/TextField";
import Check from "@material-ui/icons/Check";
import Clear from "@material-ui/icons/Clear";
import Edit from "@material-ui/icons/Edit";
import { Typography, Paper } from "@material-ui/core";

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
      width: 2,
      cellStyle: { textAlign: "left", fontWeight: "bold", maxWidth: 2 },
      headerStyle: {width: 2},
      render: props.displayInverse?(rowData) => <div>{-1*(rowData.ID-12)}</div>:null,
    },
    {
      title: "Torque",
      field: "Torque",
      type: "numeric",
      width: 15,
      cellStyle: { textAlign: "left" },
    },
    {
      title: "AD",
      field: "AD",
      type: "numeric",
      width: 15,
      cellStyle: { textAlign: "left" },
    },
  ];
  return (
    <Paper>
      <Typography variant="h5" color="primary" gutterBottom style={{padding:"5px"}}>
        {props.title}
      </Typography>
      <MaterialTable
        components={{
          EditField: (props) => (
            <TextField
              style={{ alignContent: "right", maxWidth: "120px" }}
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
          toolbar: false,
          showFirstLastPageButtons: false,
          actionsColumnIndex: -1,
          headerStyle: { textAlign: "left", fontWeight: "bold", width:2, maxWidth:2},
        }}
        icons={tableIcons}
        localization={{ pagination: { labelRowsPerPage: "12" } }}
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
