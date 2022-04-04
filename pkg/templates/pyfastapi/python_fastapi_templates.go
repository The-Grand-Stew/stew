package pyfastapi

const PyFastAPISchema string = `from typing import Optional

from pydantic import BaseModel


# Shared properties
class {{ . }}Base(BaseModel):
    pass


# Properties to receive via API on creation
class {{ . }}Create({{ . }}Base):
    pass


# Properties to receive via API on update
class {{ . }}Update({{ . }}Base):
    pass


class {{ . }}InDBBase({{ . }}Base):
    id: Optional[int] = None


# Additional properties to return via API
class {{ . }}({{ . }}InDBBase):
    pass


# Additional properties stored in DB e.g hashed password
class {{ . }}InDB({{ . }}InDBBase):
    pass`
