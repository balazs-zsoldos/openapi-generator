//
// AdditionalPropertiesNumber.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation


open class AdditionalPropertiesNumber: JSONEncodable {

    public var name: String?

    public var additionalProperties: [AnyHashable:Double] = [:]

    public init() {}

    public subscript(key: AnyHashable) -> Double? {
        get {
            if let value = additionalProperties[key] {
                return value
            }
            return nil
        }

        set {
            additionalProperties[key] = newValue
        }
    }
    // MARK: JSONEncodable
    open func encodeToJSON() -> Any {
        var nillableDictionary = [String:Any?]()
        nillableDictionary["name"] = self.name

        for (key, value) in additionalProperties {
            if let key = key as? String {
               nillableDictionary[key] = value
            }
        }

        let dictionary: [String:Any] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}

