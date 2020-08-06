import {cleanFormat} from "./Profile";

it("returnsCleanData", ()=>{
    const result = cleanFormat({
        ID: 1,
        AD: "33",
        Torque: "45"
    })
    console.log(result)
    expect(result.ID).toBe(0)
})