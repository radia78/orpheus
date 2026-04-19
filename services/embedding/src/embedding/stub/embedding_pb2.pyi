from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TextEmbeddingRequest(_message.Message):
    __slots__ = ("texts",)
    TEXTS_FIELD_NUMBER: _ClassVar[int]
    texts: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, texts: _Optional[_Iterable[str]] = ...) -> None: ...

class ImageEmbeddingRequest(_message.Message):
    __slots__ = ("images",)
    IMAGES_FIELD_NUMBER: _ClassVar[int]
    images: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, images: _Optional[_Iterable[str]] = ...) -> None: ...

class EmbeddingResponse(_message.Message):
    __slots__ = ("vectors",)
    VECTORS_FIELD_NUMBER: _ClassVar[int]
    vectors: _containers.RepeatedCompositeFieldContainer[Vector]
    def __init__(
        self, vectors: _Optional[_Iterable[_Union[Vector, _Mapping]]] = ...
    ) -> None: ...

class Vector(_message.Message):
    __slots__ = ("values",)
    VALUES_FIELD_NUMBER: _ClassVar[int]
    values: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, values: _Optional[_Iterable[float]] = ...) -> None: ...
