//
// AdditionalPropertiesClass.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation



public struct AdditionalPropertiesClass: Codable {

    public var mapString: [String:String]?
    public var mapNumber: [String:Double]?
    public var mapInteger: [String:Int]?
    public var mapBoolean: [String:Bool]?
    public var mapArrayInteger: [String:[Int]]?
    public var mapArrayAnytype: [String:[Any]]?
    public var mapMapString: [String:[String:String]]?
    public var mapMapAnytype: [String:[String:Any]]?
    public var anytype1: Any?
    public var anytype2: Any?
    public var anytype3: Any?

    public init(mapString: [String:String]?, mapNumber: [String:Double]?, mapInteger: [String:Int]?, mapBoolean: [String:Bool]?, mapArrayInteger: [String:[Int]]?, mapArrayAnytype: [String:[Any]]?, mapMapString: [String:[String:String]]?, mapMapAnytype: [String:[String:Any]]?, anytype1: Any?, anytype2: Any?, anytype3: Any?) {
        self.mapString = mapString
        self.mapNumber = mapNumber
        self.mapInteger = mapInteger
        self.mapBoolean = mapBoolean
        self.mapArrayInteger = mapArrayInteger
        self.mapArrayAnytype = mapArrayAnytype
        self.mapMapString = mapMapString
        self.mapMapAnytype = mapMapAnytype
        self.anytype1 = anytype1
        self.anytype2 = anytype2
        self.anytype3 = anytype3
    }

    public enum CodingKeys: String, CodingKey { 
        case mapString = "map_string"
        case mapNumber = "map_number"
        case mapInteger = "map_integer"
        case mapBoolean = "map_boolean"
        case mapArrayInteger = "map_array_integer"
        case mapArrayAnytype = "map_array_anytype"
        case mapMapString = "map_map_string"
        case mapMapAnytype = "map_map_anytype"
        case anytype1 = "anytype_1"
        case anytype2 = "anytype_2"
        case anytype3 = "anytype_3"
    }


}

