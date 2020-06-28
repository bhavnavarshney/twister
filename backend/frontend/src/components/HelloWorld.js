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

function cleanFormat(rowData) {
  return {
    ID: rowData.ID,
    AD: parseInt(rowData.AD),
    Torque: parseInt(rowData.Torque),
  };
}

export default function HelloWorld() {
  const [profile, setProfile] = React.useState([]);

  const handleRead = () => {
    window.backend.Drill.GetProfile().then((result) => {
      const newProfile = mapFieldsToProfile(result.Fields);
      setProfile(newProfile);
    });
  };

  const handleWrite = () => {
    const cleanProfile = profile.map((row) => cleanFormat(row));
    window.backend.Drill.WriteProfile(cleanProfile).then((result) => {
      console.log(result);
    });
  };

  const rowUpdateHandler = (newData, oldData) =>
    new Promise((resolve) => {
      setTimeout(() => {
        resolve();
        if (oldData) {
          const data = [...profile];
          data[data.indexOf(oldData)] = cleanFormat(newData);
          console.log(cleanFormat(newData));
          setProfile(data);
          window.backend.Drill.WriteParam(cleanFormat(newData)).then(
            (result) => {
              console.log(result);
            }
          );
        }
      }, 600);
    });

  return (
    <div className="App">
      <Grid container spacing={3}>
        <Grid item xs={2} style={{ minWidth: "300px" }}>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <InfoCard />
            </Grid>
          </Grid>
        </Grid>
        <Grid
          item
          xs={4}
          style={{
            minWidth: "300px",
          }}
        >
          <ParamTable
            id="unique"
            title="Clockwise"
            params={profile.slice(0, 12)}
            handleRowUpdate={rowUpdateHandler}
          />
        </Grid>
        <Grid
          item
          xs={4}
          style={{
            minWidth: "300px",
          }}
        >
          <ParamTable
            title="Counterclockwise"
            params={profile.slice(12, 24)}
            handleRowUpdate={rowUpdateHandler}
          />
        </Grid>
        <Grid item xs={1}>
          <Paper>
            <Grid container spacing={0}>
              <Grid item xs={12}>
                <Button
                  onClick={handleRead}
                  variant="contained"
                  color="primary"
                >
                  Read
                </Button>
              </Grid>
            </Grid>
          </Paper>
        </Grid>
      </Grid>
    </div>
  );
}
