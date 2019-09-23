//
// TypeHolderExample.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation


open class TypeHolderExample: JSONEncodable {

    public var stringItem: String?
    public var numberItem: Double?
    public var floatItem: Float?
    public var integerItem: Int32?
    public var boolItem: Bool?
    public var arrayItem: [Int32]?

    public init() {}

    // MARK: JSONEncodable
    open func encodeToJSON() -> Any {
        var nillableDictionary = [String:Any?]()
        nillableDictionary["string_item"] = self.stringItem
        nillableDictionary["number_item"] = self.numberItem
        nillableDictionary["float_item"] = self.floatItem
        nillableDictionary["integer_item"] = self.integerItem?.encodeToJSON()
        nillableDictionary["bool_item"] = self.boolItem
        nillableDictionary["array_item"] = self.arrayItem?.encodeToJSON()

        let dictionary: [String:Any] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}
