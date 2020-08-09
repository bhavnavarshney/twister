export function mapFieldsToProfile(fields) {
    if (fields === null) return []
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