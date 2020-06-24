import React, { forwardRef } from 'react';
import MaterialTable from 'material-table';
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

export default function ParamTable() {
    const [state, setState] = React.useState({
        columns: [
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
          ],
        data: [
            { ID: 1, Torque: 15, AD: 15 },
            { ID: 2, Torque: 15, AD: 15 },
            { ID: 3, Torque: 15, AD: 15 },
            { ID: 4, Torque: 15, AD: 15 },
            { ID: 5, Torque: 15, AD: 15 },
            { ID: 6, Torque: 15, AD: 15 },
            { ID: 7, Torque: 15, AD: 15 },
            { ID: 8, Torque: 15, AD: 15 },
            { ID: 9, Torque: 15, AD: 15 },
            { ID: 10, Torque: 15, AD: 15 },
            { ID: 11, Torque: 15, AD: 15 },
            { ID: 12, Torque: 15, AD: 15 },
        ],
      });
return (<Paper>
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
      columns={state.columns}
      data={state.data}
      editable={{
        isEditable: (rowData) => rowData.name !== "ID",
        isDeleteHidden: (rowData) => true,
        onRowUpdate: (newData, oldData) =>
          new Promise((resolve) => {
            setTimeout(() => {
              resolve();
              if (oldData) {
                setState((prevState) => {
                    const data = [...prevState.data];
                    data[data.indexOf(oldData)] = newData;
                    return { ...prevState, data };
                  });
              }
            }, 600);
          }),
      }}
    />
  </Paper>);
}

