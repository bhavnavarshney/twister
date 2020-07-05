import React, {useEffect} from "react";
import { useSnackbar } from 'notistack';
import Grid from "@material-ui/core/Grid";
import InfoCard from "./InfoCard";
import ParamTable from "./ParamTable";
import Wails from "@wailsapp/runtime"

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
  const [info, setInfo] = React.useState({});
  const [currentOffset, setCurrentOffset] = React.useState(null);
  const [port, setPort] = React.useState(3);
  const [profile, setProfile] = React.useState([]);
  const [isConnected, setIsConnected] = React.useState(false);
  const { enqueueSnackbar } = useSnackbar()

  const infoSnackBarOptions = {variant: "info", autoHideDuration:3000, anchorOrigin:{
    vertical: "bottom",
    horizontal: "right",
  }}
  
  const errorSnackBarOptions = {variant: "error", anchorOrigin:{
    vertical: "bottom",
    horizontal: "right",
  }}

  const successSnackBarOptions = {variant: "success", autoHideDuration:3000, anchorOrigin:{
    vertical: "bottom",
    horizontal: "right",
  }}

  useEffect(() => {
    Wails.Events.On("CurrentOffset", message => {
      setCurrentOffset(message)
    });
  }, []);


  const handleClose = () => {
    window.backend.Drill.Close().then((result) => {
      setCurrentOffset(null)
      setInfo({})
      setProfile([])
      enqueueSnackbar("Closed", infoSnackBarOptions)
      setIsConnected(false)
    }).catch((err)=> {
      enqueueSnackbar("Error Closing port:" + err, infoSnackBarOptions)
      setIsConnected(false)
    });
  }

  const handleSetPort = (e) => {
    setPort(e.target.value)
  }
  const handleRead = () => {
    window.backend.Drill.Open("COM" + port.toString()).then((result)=>{
      setIsConnected(true)
      enqueueSnackbar("Drill Connected", successSnackBarOptions)
      window.backend.Drill.GetInfo().then((result) => {
        setInfo(result)
        setCurrentOffset(result.CurrentOffset)
        window.backend.Drill.GetProfile().then((result) => {
          const newProfile = mapFieldsToProfile(result.Fields);
          setProfile(newProfile);
          setIsConnected(true)
        });
      }).catch((err)=>{
        console.log(err)
        enqueueSnackbar("Error getting info" + err, errorSnackBarOptions);
        setIsConnected(false)
      });
    }).catch((err)=>{
      enqueueSnackbar("Error connecting: "+err, errorSnackBarOptions);
      setIsConnected(false)
    })
    

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
          window.backend.Drill.WriteParam(cleanFormat(newData)).then(
            (result) => {
              enqueueSnackbar("Parameter Saved", successSnackBarOptions);
            }
          ).catch((err)=>{
            enqueueSnackbar("Error saving:" + err, errorSnackBarOptions);
          });
        }
      }, 600);
    });

  return (
    <div className="App">
      <Grid container spacing={3}>
        <Grid item xs={2} style={{ minWidth: "300px" }}>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <InfoCard isConnected={isConnected} data={info} currentOffset={currentOffset} handleOpen={handleRead} handleClose={handleClose} handleSetPort={handleSetPort}/>
            </Grid>
          </Grid>
        </Grid>
        <Grid
          item
          xs={4}
          style={{
            minWidth: "400px",
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
            minWidth: "400px",
          }}
        >
          <ParamTable
            title="Counterclockwise"
            params={profile.slice(12, 24)}
            handleRowUpdate={rowUpdateHandler}
          />
        </Grid>
      </Grid>
    </div>
  );
}
