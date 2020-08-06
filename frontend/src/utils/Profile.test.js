import {cleanFormat, mapFieldsToProfile} from "./Profile";

it("returnsCleanData", ()=>{
    const result = cleanFormat({
        ID: 1,
        AD: "33",
        Torque: "45"
    })
    console.log(result)
    expect(result.ID).toBe(0)
})

it("mapsFieldsToProfile", ()=>{
    const fields =  [{
            Torque: 30,
            AD: 30,
        },]
    const result = mapFieldsToProfile(fields)
    expect(result).toEqual([{ID: 1, Torque: 30, AD: 30}])
})

it("mapsFieldsToProfileHandlesNull", ()=>{
    const result = mapFieldsToProfile(null)
    expect(result).toEqual([])
})