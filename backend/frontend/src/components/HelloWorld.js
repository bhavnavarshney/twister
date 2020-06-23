import React, { forwardRef } from "react";
import Modal from "react-modal";
import MaterialTable from "material-table";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import AddBox from "@material-ui/icons/AddBox";
import ArrowDownward from "@material-ui/icons/ArrowDownward";
import Check from "@material-ui/icons/Check";
import ChevronLeft from "@material-ui/icons/ChevronLeft";
import ChevronRight from "@material-ui/icons/ChevronRight";
import Clear from "@material-ui/icons/Clear";
import DeleteOutline from "@material-ui/icons/DeleteOutline";
import Edit from "@material-ui/icons/Edit";
import FilterList from "@material-ui/icons/FilterList";
import FirstPage from "@material-ui/icons/FirstPage";
import LastPage from "@material-ui/icons/LastPage";
import Remove from "@material-ui/icons/Remove";
import SaveAlt from "@material-ui/icons/SaveAlt";
import Search from "@material-ui/icons/Search";
import ViewColumn from "@material-ui/icons/ViewColumn";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    padding: theme.spacing(2),
    textAlign: "center",
    color: theme.palette.text.secondary,
  },
}));

const tableIcons = {
  Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
  Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
  Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
  Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref} />),
  DetailPanel: forwardRef((props, ref) => (
    <ChevronRight {...props} ref={ref} />
  )),
  Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
  Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref} />),
  Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref} />),
  FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
  LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
  NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
  PreviousPage: forwardRef((props, ref) => (
    <ChevronLeft {...props} ref={ref} />
  )),
  ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
  Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
  SortArrow: forwardRef((props, ref) => <ArrowDownward {...props} ref={ref} />),
  ThirdStateCheck: forwardRef((props, ref) => <Remove {...props} ref={ref} />),
  ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref} />),
};

class HelloWorld extends React.Component {
  constructor(props, context) {
    super();
    this.state = {
      showModal: false,
      columns: [
        { title: "ID", field: "ID", type: "numeric", editable: "never" },
        { title: "Torque", field: "Torque", type: "numeric" },
        { title: "AD", field: "AD", type: "numeric" },
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
      dataCW: [
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
      dataCCW: [
        { ID: 13, Torque: 15, AD: 15 },
        { ID: 14, Torque: 15, AD: 15 },
        { ID: 15, Torque: 15, AD: 15 },
        { ID: 16, Torque: 15, AD: 15 },
        { ID: 17, Torque: 15, AD: 15 },
        { ID: 18, Torque: 15, AD: 15 },
        { ID: 19, Torque: 15, AD: 15 },
        { ID: 20, Torque: 15, AD: 15 },
        { ID: 21, Torque: 15, AD: 15 },
        { ID: 22, Torque: 15, AD: 15 },
        { ID: 23, Torque: 15, AD: 15 },
        { ID: 24, Torque: 15, AD: 15 },
      ],
    };

    this.handleOpenModal = this.handleOpenModal.bind(this);
    this.handleCloseModal = this.handleCloseModal.bind(this);
  }

  handleOpenModal() {
    this.setState({ showModal: true });

    window.backend.basic().then((result) =>
      this.setState({
        result,
      })
    );
  }

  handleCloseModal() {
    this.setState({ showModal: false });
  }

  render() {
    const { result } = this.state;
    console.log(result);
    return (
      <div className="App">
        <button onClick={this.handleOpenModal} type="button">
          Hello
        </button>
        

        <Grid container spacing={3}>
          <Grid item xs={3}>
          <Paper>
          <Grid container spacing={3}>
         
          <Grid item xs={12}>
            <Button variant="contained" color="primary">
              Open
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Button variant="contained" color="primary">
              Read
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Button variant="contained" color="primary">
              Send
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Button variant="contained" color="secondary">
              Close
            </Button>
          </Grid>
         
        </Grid>
        </Paper>
            </Grid>
          <Grid item xs={3}>
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
                columns={this.state.columns}
                data={this.state.dataCW}
                editable={{
                  isEditable: (rowData) => rowData.name !== "ID",
                  isDeleteHidden: (rowData) => true,
                  onRowAdd: (newData) =>
                    new Promise((resolve) => {
                      setTimeout(() => {
                        resolve();
                        const data = [...this.state.data];
                        data.push(newData);
                        this.setState({ ...this.state, data });
                      }, 600);
                    }),
                  onRowUpdate: (newData, oldData) =>
                    new Promise((resolve) => {
                      setTimeout(() => {
                        resolve();
                        if (oldData) {
                          const data = [...this.state.data];
                          data[data.indexOf(oldData)] = newData;
                          this.setState({ ...this.state, data });
                        }
                      }, 600);
                    }),
                  onRowDelete: (oldData) =>
                    new Promise((resolve) => {
                      setTimeout(() => {
                        resolve();
                        const data = [...this.state.data];
                        data.splice(data.indexOf(oldData), 1);
                        this.setState({ ...this.state, data });
                      }, 600);
                    }),
                }}
              />
            </Paper>
          </Grid>
          <Grid item xs={3}>
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
                columns={this.state.columns}
                data={this.state.dataCCW}
                editable={{
                  isDeleteHidden: (rowData) => true,
                  onRowAdd: (newData) =>
                    new Promise((resolve) => {
                      setTimeout(() => {
                        resolve();
                        const data = [...this.state.data];
                        data.push(newData);
                        this.setState({ ...this.state, data });
                      }, 600);
                    }),
                  onRowUpdate: (newData, oldData) =>
                    new Promise((resolve) => {
                      setTimeout(() => {
                        resolve();
                        if (oldData) {
                          const data = [...this.state.data];
                          data[data.indexOf(oldData)] = newData;
                          this.setState({ ...this.state, data });
                        }
                      }, 600);
                    }),
                  onRowDelete: (oldData) =>
                    new Promise((resolve) => {
                      setTimeout(() => {
                        resolve();
                        const data = [...this.state.data];
                        data.splice(data.indexOf(oldData), 1);
                        this.setState({ ...this.state, data });
                      }, 600);
                    }),
                }}
              />
            </Paper>
          </Grid>
        </Grid>

        <Modal
          isOpen={this.state.showModal}
          contentLabel="Minimal Modal Example"
        >
          {!result
            ? null
            : result.Fields.map((i) => (
                <p>
                  <p>{i.AD}</p>
                  <p>{i.Torque}</p>
                </p>
              ))}

          <button onClick={this.handleCloseModal}>Close Modal</button>
        </Modal>
      </div>
    );
  }
}

export default HelloWorld;
