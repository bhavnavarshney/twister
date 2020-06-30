import React, {useEffect} from "react";
import Grid from "@material-ui/core/Grid";
import InfoCard from "./InfoCard";
import ParamTable from "./ParamTable";
import SnackBar from "./SnackBar";
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
  const [showSnackBar, setShowSnackBar] = React.useState({message: "", severity: "info"});
  const [info, setInfo] = React.useState({});
  const [currentOffset, setCurrentOffset] = React.useState(null);
  const [port, setPort] = React.useState(3);
  const [profile, setProfile] = React.useState([]);

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
      setShowSnackBar({message: "Closed", severity: "success"})
      console.log(result)
    }).catch((err)=> {
      console.log(err)
    });
  }
  const handleCloseSnackBar = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setShowSnackBar({message: "", severity: "info"});
  };

  const handleSetPort = (e) => {
    setPort(e.target.value)
  }
  const handleRead = () => {
    window.backend.Drill.Open("COM" + port.toString()).then((result)=>{
     setShowSnackBar({message: "Drill Connected", severity: "success"})
      window.backend.Drill.GetInfo().then((result) => {
        setInfo(result)
        setCurrentOffset(result.CurrentOffset)
        window.backend.Drill.GetProfile().then((result) => {
          const newProfile = mapFieldsToProfile(result.Fields);
          setProfile(newProfile);
        });
      }).catch((err)=>{
        console.log(err)
        setShowSnackBar({message: "Error getting info" + err, severity: "error"})
      });
    }).catch((err)=>{
      setShowSnackBar({message: "Error connecting: "+err, severity: "error"})
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
              setShowSnackBar({message: "Parameter Saved", severity: "success"})
            }
          ).catch((err)=>{
            setShowSnackBar({message: "Error saving:" + err, severity: "error"})
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
              <InfoCard data={info} currentOffset={currentOffset} handleOpen={handleRead} handleClose={handleClose} handleSetPort={handleSetPort}/>
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
      <SnackBar message={showSnackBar.message} severity={showSnackBar.severity} handleClose = {handleCloseSnackBar}/>
    </div>
  );
}
