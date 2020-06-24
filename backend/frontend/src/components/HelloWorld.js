import React from "react";
import Modal from "react-modal";
import Button from "@material-ui/core/Button";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import InfoCard from "./InfoCard";
import ParamTable from "./ParamTable";

class HelloWorld extends React.Component {
  constructor(props, context) {
    super();
    this.state = {
      showModal: false,
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
        <Grid container spacing={3}>
          <Grid item xs={3}>
            <Grid item xs={12}>
              <InfoCard />
            </Grid>
            <Grid item xs={12}>
              <Paper>
                <Grid container spacing={3}>
                  <Grid item xs={12}>
                    <Button
                      onClick={this.handleOpenModal}
                      variant="contained"
                      color="primary"
                    >
                      Read
                    </Button>
                  </Grid>
                  <Grid item xs={12}>
                    <Button variant="contained" color="primary">
                      Write
                    </Button>
                  </Grid>
                </Grid>
              </Paper>
            </Grid>
          </Grid>
          <Grid item xs={3}>
            <ParamTable />
          </Grid>
          <Grid item xs={3}>
            <ParamTable />
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
