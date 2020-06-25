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
  const [profile, setProfile] = React.useState([]);

  const handleOpenModal = () => {
    window.backend.basic().then((result) => {
      const newProfile = mapFieldsToProfile(result.Fields);
      setProfile(newProfile);
    });
  };

  const rowUpdateHandler = (newData, oldData) =>
    new Promise((resolve) => {
      setTimeout(() => {
        resolve();
        if (oldData) {
          const data = [...profile];
          data[data.indexOf(oldData)] = newData;
          setProfile(data);
        }
      }, 600);
    });

  return (
    <div className="App">
      <Grid container spacing={3}>
        <Grid item xs={3}>
          <Grid container spacing={3}>
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
        </Grid>
        <Grid item xs={3}>
          <ParamTable
            id="unique"
            params={profile.slice(0,12)}
            handleRowUpdate={rowUpdateHandler}
          />
        </Grid>
        <Grid item xs={3}>
          <ParamTable
            params={profile.slice(12,24)}
            handleRowUpdate={rowUpdateHandler}
          />
        </Grid>
      </Grid>
    </div>
  );
}
