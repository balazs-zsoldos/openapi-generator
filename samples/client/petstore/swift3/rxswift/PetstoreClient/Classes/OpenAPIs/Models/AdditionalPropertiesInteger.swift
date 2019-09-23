//
// AdditionalPropertiesInteger.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation


open class AdditionalPropertiesInteger: JSONEncodable {

    public var name: String?

    public var additionalProperties: [AnyHashable:Int32] = [:]

    public init() {}

    public subscript(key: AnyHashable) -> Int32? {
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

