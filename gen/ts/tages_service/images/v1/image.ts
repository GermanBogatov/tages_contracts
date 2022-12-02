// @generated by protobuf-ts 2.2.2 with parameter long_type_string
// @generated from protobuf file "tages_service/images/v1/image.proto" (package "tages_service.images.v1", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message tages_service.images.v1.Image
 */
export interface Image {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: bytes image = 2;
     */
    image: Uint8Array;
    /**
     * @generated from protobuf field: int64 create_at = 3;
     */
    createAt: string;
    /**
     * @generated from protobuf field: int64 updated_at = 4;
     */
    updatedAt: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class Image$Type extends MessageType<Image> {
    constructor() {
        super("tages_service.images.v1.Image", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "image", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 3, name: "create_at", kind: "scalar", T: 3 /*ScalarType.INT64*/ },
            { no: 4, name: "updated_at", kind: "scalar", T: 3 /*ScalarType.INT64*/ }
        ]);
    }
    create(value?: PartialMessage<Image>): Image {
        const message = { name: "", image: new Uint8Array(0), createAt: "0", updatedAt: "0" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Image>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Image): Image {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* bytes image */ 2:
                    message.image = reader.bytes();
                    break;
                case /* int64 create_at */ 3:
                    message.createAt = reader.int64().toString();
                    break;
                case /* int64 updated_at */ 4:
                    message.updatedAt = reader.int64().toString();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Image, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* bytes image = 2; */
        if (message.image.length)
            writer.tag(2, WireType.LengthDelimited).bytes(message.image);
        /* int64 create_at = 3; */
        if (message.createAt !== "0")
            writer.tag(3, WireType.Varint).int64(message.createAt);
        /* int64 updated_at = 4; */
        if (message.updatedAt !== "0")
            writer.tag(4, WireType.Varint).int64(message.updatedAt);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message tages_service.images.v1.Image
 */
export const Image = new Image$Type();
