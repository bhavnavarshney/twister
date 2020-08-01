import React from "react";
import { withStyles, makeStyles } from "@material-ui/core/styles";
import Tooltip from "@material-ui/core/Tooltip";
import Button from "@material-ui/core/Button";
import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Switch from "@material-ui/core/Switch";
import Grid from "@material-ui/core/Grid";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import InputAdornment from "@material-ui/core/InputAdornment";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";

const StyledLabelTableCell = withStyles((theme) => ({
  body: {
    fontWeight: "bold",
  },
}))(TableCell);

const StyledDataTableCell = withStyles((theme) => ({
  body: {
    textAlign: "right",
  },
}))(TableCell);

const useStyles = makeStyles({
  root: {
    minWidth: 275,
  },
  bullet: {
    display: "inline-block",
    margin: "0 2px",
    transform: "scale(0.8)",
  },
  title: {
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
  table: {},
});

function createData(name, field) {
  return { name, field };
}

export default function InfoCard(props) {
  const classes = useStyles();
  const rows = [
    createData("Drill ID", props.data.DrillID),
    createData("Drill Type", props.data.DrillType),
    createData("Calibrated Offset", props.data.CalibratedOffset),
    createData("Current Offset", props.currentOffset),
    createData("Status", !props.data.DrillID ? "Not Connected" : "Connected"),
  ];

  const handleSwitch = () => {
    if (props.isConnected) {
      props.handleClose();
    } else {
      props.handleOpen();
    }
  };

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography
          className={classes.title}
          color="textSecondary"
          gutterBottom
        >
          Drill
        </Typography>
        <Tooltip placement="right" title="Find this using Device Manager">
          <TextField
            onChange={props.handleSetPort}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">COM</InputAdornment>
              ),
              min: "0",
              max: "65535",
              step: "1",
            }}
            type="number"
            style={{ padding: "15px" }}
            label="Device Port"
            color="secondary"
            defaultValue="3"
          />
        </Tooltip>

        <Table className={classes.table} aria-label="simple table">
          <TableBody>
            {rows.map((row) => (
              <TableRow key={row.name}>
                <StyledLabelTableCell component="th" scope="row">
                  {row.name}
                </StyledLabelTableCell>
                <StyledDataTableCell>{row.field}</StyledDataTableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </CardContent>
      <CardActions>
        <Grid container spacing={1}>
          <Grid item xs={12}>
            <FormControlLabel
              control={
                <Switch
                  checked={props.isConnected}
                  onChange={handleSwitch}
                  color="primary"
                />
              }
              label="Connect"
            />
          </Grid>
          <Grid item xs={12}>
            <Button
              variant="contained"
              color="primary"
              disabled={!props.data.DrillID}
              onClick={props.handleGetCurrentOffset}
            >
              Current Offset
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Button
              variant="contained"
              color="primary"
              disabled={!props.data.DrillID}
              onClick={props.handleSave}
            >
              Save
            </Button>
          </Grid>
          <Grid item xs={12}>
            <div>
              Load
              <input type="file" id="input" accept=".csv" onChange={ (e) => props.handleLoad(e.target.files) }></input>
              </div>
          </Grid>
          <Grid item xs={12}>
            <Button
              variant="contained"
              color="primary"
              disabled={!props.data.DrillID}
              
            >
              Finish
            </Button>
          </Grid>
        </Grid>
      </CardActions>
    </Card>
  );
}
