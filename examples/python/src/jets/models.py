# Code generated by sqlc. DO NOT EDIT.
import dataclasses


@dataclasses.dataclass()
class Jet:
    id: int
    pilot_id: int
    age: int
    name: str
    color: str


@dataclasses.dataclass()
class Language:
    id: int
    language: str


@dataclasses.dataclass()
class Pilot:
    id: int
    name: str


@dataclasses.dataclass()
class PilotLanguage:
    pilot_id: int
    language_id: int
