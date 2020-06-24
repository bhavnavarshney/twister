import React from "react";
import Button from "@material-ui/core/Button";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import InfoCard from "./InfoCard";
import ParamTable from "./ParamTable";

function mapFieldsToProfile(fields) {
  return fields.map((item, index) => {
    return {
      ID: index + 1,
      Torque: item.Torque,
      AD: item.AD,
    };
  });
}

export default function HelloWorld() {
  const [state, setState] = React.useState({
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
    profile: [],
  });

  const handleOpenModal = () => {
    window.backend.basic().then((result) => {
      var newState = {...state};
      newState.profile = mapFieldsToProfile(result.Fields);
      setState(newState);
    });
  };

  const rowUpdateHandler = (newData, oldData) =>
    new Promise((resolve) => {
      setTimeout(() => {
        resolve();
        if (oldData) {
          const data = [...state.profile];
          data[data.indexOf(oldData)] = newData;
          setState({...state, data});
        }
      }, 600);
    });

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
                    onClick={handleOpenModal}
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
          <ParamTable
            id="unique"
            params={state.profile}
            handleRowUpdate={rowUpdateHandler}
          />
        </Grid>
        <Grid item xs={3}>
          <ParamTable
            params={state.profile}
            handleRowUpdate={rowUpdateHandler}
          />
        </Grid>
      </Grid>
    </div>
  );
}
