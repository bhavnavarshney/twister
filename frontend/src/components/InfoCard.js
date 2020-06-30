import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Tooltip from "@material-ui/core/Tooltip";
import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import InputAdornment from "@material-ui/core/InputAdornment";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";

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
    createData("Current Offset", props.currentOffset ),
    createData("Status", !props.data.DrillID? "Not Connected":"Connected"),
  ];

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
                <TableCell component="th" scope="row">
                  {row.name}
                </TableCell>
                <TableCell align="right">{row.field}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </CardContent>
      <CardActions>
        <Button variant="contained" color="primary" onClick={props.handleOpen}>
          Open
        </Button>
        <Button
          variant="contained"
          color="secondary"
          onClick={props.handleClose}
        >
          Close
        </Button>
      </CardActions>
    </Card>
  );
}
