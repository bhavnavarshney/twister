import React from "react";
import { useSnackbar } from "notistack";
import Grid from "@material-ui/core/Grid";
import InfoCard from "./InfoCard";
import ParamTable from "./ParamTable";

export function mapFieldsToProfile(fields) {
  return fields.map((item, index) => {
    return {
      ID: index + 1,
      Torque: item.Torque,
      AD: item.AD,
    };
  });
}

// cleanFormat converts the data from string to integer
// It also removes the offset on the ID, so that 1-24 is mapped to 0-23
export function cleanFormat(rowData) {
  return {
    ID: rowData.ID - 1,
    AD: parseInt(rowData.AD),
    Torque: parseInt(rowData.Torque),
  };
}

export default function ParamView() {
  const [info, setInfo] = React.useState({});
  const [currentOffset, setCurrentOffset] = React.useState(null);
  const [port, setPort] = React.useState(3);
  const [profile, setProfile] = React.useState([]);
  const [isConnected, setIsConnected] = React.useState(false);
  const { enqueueSnackbar } = useSnackbar();

  const infoSnackBarOptions = {
    variant: "info",
    autoHideDuration: 3000,
    anchorOrigin: {
      vertical: "bottom",
      horizontal: "right",
    },
  };

  const errorSnackBarOptions = {
    variant: "error",
    anchorOrigin: {
      vertical: "bottom",
      horizontal: "right",
    },
  };

  const successSnackBarOptions = {
    variant: "success",
    autoHideDuration: 3000,
    anchorOrigin: {
      vertical: "bottom",
      horizontal: "right",
    },
  };

  // useEffect(() => {
  //   Wails.Events.On("CurrentOffset", (message) => {
  //     setCurrentOffset(message);
  //   });
  // }, []);

  const handleLoad = (fileList) => {
    /*global LoadProfile*/
    /*eslint no-undef: "error"*/
    console.log(fileList[0]);
    fileList[0].text().then((fileContent) => {
      LoadProfile(fileContent)
        .then((result) => {
          console.log(result);
          const newProfile = mapFieldsToProfile(result.Fields);
          setProfile(newProfile);
          enqueueSnackbar("Profile Loaded!", successSnackBarOptions);
        })
        .catch((err) => {
          enqueueSnackbar("Error loading profile:" + err, errorSnackBarOptions);
        });
    });
  };

  const handleSave = () => {
    /*global SaveProfile*/
    /*eslint no-undef: "error"*/
    SaveProfile("./demo.csv")
      .then((result) => {
        enqueueSnackbar("Profile Saved!", successSnackBarOptions);
      })
      .catch((err) => {
        enqueueSnackbar("Error saving profile:" + err, errorSnackBarOptions);
      });
  };

  const handleClose = () => {
    /*global Close*/
    /*eslint no-undef: "error"*/
    Close()
      .then((result) => {
        setCurrentOffset(null);
        setInfo({});
        setProfile([]);
        enqueueSnackbar("Closed", infoSnackBarOptions);
        setIsConnected(false);
      })
      .catch((err) => {
        enqueueSnackbar("Error Closing port:" + err, infoSnackBarOptions);
        setIsConnected(false);
      });
  };

  const handleGetCurrentOffset = () => {
    /*global GetCurrentOffset*/
    /*eslint no-undef: "error"*/
    GetCurrentOffset()
      .then((result) => {
        setCurrentOffset(result);
        enqueueSnackbar("Current Offset Received", successSnackBarOptions);
      })
      .catch((err) => {
        enqueueSnackbar(
          "Error reading offset. Please try again.",
          errorSnackBarOptions
        );
      });
  };

  const handleSetPort = (e) => {
    setPort(e.target.value);
  };
  const handleRead = () => {
    /*global Open*/
    /*eslint no-undef: "error"*/
    Open("COM" + port.toString())
      .then((result) => {
        setIsConnected(true);
        enqueueSnackbar("Drill Connected", successSnackBarOptions);
        /*global GetInfo*/
        /*eslint no-undef: "error"*/
        GetInfo()
          .then((result) => {
            setInfo(result);
            setCurrentOffset(result.CurrentOffset);
            /*global GetProfile*/
            /*eslint no-undef: "error"*/
            GetProfile().then((result) => {
              const newProfile = mapFieldsToProfile(result.Fields);
              setProfile(newProfile);
              setIsConnected(true);
            });
          })
          .catch((err) => {
            console.log(err);
            enqueueSnackbar("Error getting info" + err, errorSnackBarOptions);
            setIsConnected(false);
          });
      })
      .catch((err) => {
        enqueueSnackbar("Error connecting: " + err, errorSnackBarOptions);
        setIsConnected(false);
      });
  };

  // const handleWrite = () => {
  //   const cleanProfile = profile.map((row) => cleanFormat(row));
  //   window.backend.Drill.WriteProfile(cleanProfile).then((result) => {
  //     console.log(result);
  //   });
  // };

  const rowUpdateHandler = (newData, oldData) =>
    new Promise((resolve) => {
      setTimeout(() => {
        resolve();
        if (oldData) {
          const data = [...profile];
          data[data.indexOf(oldData)] = cleanFormat(newData);
          setProfile(data);
          /*global WriteParam*/
          /*eslint no-undef: "error"*/
          WriteParam(cleanFormat(newData))
            .then((result) => {
              enqueueSnackbar("Parameter Saved", successSnackBarOptions);
            })
            .catch((err) => {
              enqueueSnackbar("Error saving:" + err, errorSnackBarOptions);
            });
        }
      }, 600);
    });

  return (
    <div className="App" style={{ height: "100%" }}>
      <Grid container spacing={1}>
        <Grid
          style={{
            minWidth: "288px",
            maxWidth: "288px",
          }}
          item
        >
          <InfoCard
            isConnected={isConnected}
            data={info}
            currentOffset={currentOffset}
            handleOpen={handleRead}
            handleClose={handleClose}
            handleSetPort={handleSetPort}
            handleGetCurrentOffset={handleGetCurrentOffset}
            handleSave={handleSave}
            handleLoad={handleLoad}
          />
        </Grid>
        <Grid
          item
          style={{
            minWidth: "457px",
            maxWidth: "457px",
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
          style={{
            minWidth: "457px",
            maxWidth: "457px",
          }}
        >
          <ParamTable
            title="Counterclockwise"
            params={profile.slice(12, 24)}
            handleRowUpdate={rowUpdateHandler}
            displayInverse={true}
          />
        </Grid>
      </Grid>
    </div>
  );
}
