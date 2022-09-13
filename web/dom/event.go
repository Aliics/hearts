package dom

type EventType string

const (
	EventTypeAbort              EventType = "abort"
	EventTypeAfterPrint         EventType = "afterprint"
	EventTypeAnimationEnd       EventType = "animationend"
	EventTypeAnimationIteration EventType = "animationiteration"
	EventTypeAnimationStart     EventType = "animationstart"
	EventTypeBeforePrint        EventType = "beforeprint"
	EventTypeBeforeUnload       EventType = "beforeunload"
	EventTypeBlur               EventType = "blur"
	EventTypeCanplay            EventType = "canplay"
	EventTypeCanPlayThrough     EventType = "canplaythrough"
	EventTypeChange             EventType = "change"
	EventTypeClick              EventType = "click"
	EventTypeContextMenu        EventType = "contextmenu"
	EventTypeCopy               EventType = "copy"
	EventTypeCut                EventType = "cut"
	EventTypeDoubleClick        EventType = "dblclick"
	EventTypeDrag               EventType = "drag"
	EventTypeDragEnd            EventType = "dragend"
	EventTypeDragEnter          EventType = "dragenter"
	EventTypeDragLeave          EventType = "dragleave"
	EventTypeDragOver           EventType = "dragover"
	EventTypeDragStart          EventType = "dragstart"
	EventTypeDrop               EventType = "drop"
	EventTypeDurationChange     EventType = "durationchange"
	EventTypeEnded              EventType = "ended"
	EventTypeError              EventType = "error"
	EventTypeFocus              EventType = "focus"
	EventTypeFocusIn            EventType = "focusin"
	EventTypeFocusOut           EventType = "focusout"
	EventTypeFullscreenChange   EventType = "fullscreenchange"
	EventTypeFullscreenError    EventType = "fullscreenerror"
	EventTypeHashChange         EventType = "hashchange"
	EventTypeInput              EventType = "input"
	EventTypeInvalid            EventType = "invalid"
	EventTypeKeyDown            EventType = "keydown"
	EventTypeKeyPress           EventType = "keypress"
	EventTypeKeyUp              EventType = "keyup"
	EventTypeLoad               EventType = "load"
	EventTypeLoadedData         EventType = "loadeddata"
	EventTypeLoadedMetadata     EventType = "loadedmetadata"
	EventTypeLoadStart          EventType = "loadstart"
	EventTypeMessage            EventType = "message"
	EventTypeMousedown          EventType = "mousedown"
	EventTypeMouseEnter         EventType = "mouseenter"
	EventTypeMouseLeave         EventType = "mouseleave"
	EventTypeMouseMove          EventType = "mousemove"
	EventTypeMouseOver          EventType = "mouseover"
	EventTypeMouseOut           EventType = "mouseout"
	EventTypeMouseUp            EventType = "mouseup"
	EventTypeOffline            EventType = "offline"
	EventTypeOnline             EventType = "online"
	EventTypeOpen               EventType = "open"
	EventTypePageHide           EventType = "pagehide"
	EventTypePageShow           EventType = "pageshow"
	EventTypePaste              EventType = "paste"
	EventTypePause              EventType = "pause"
	EventTypePlay               EventType = "play"
	EventTypePlaying            EventType = "playing"
	EventTypePopState           EventType = "popstate"
	EventTypeProgress           EventType = "progress"
	EventTypeRateChange         EventType = "ratechange"
	EventTypeResize             EventType = "resize"
	EventTypeReset              EventType = "reset"
	EventTypeScroll             EventType = "scroll"
	EventTypeSearch             EventType = "search"
	EventTypeSeeked             EventType = "seeked"
	EventTypeSeeking            EventType = "seeking"
	EventTypeSelect             EventType = "select"
	EventTypeShow               EventType = "show"
	EventTypeStalled            EventType = "stalled"
	EventTypeStorage            EventType = "storage"
	EventTypeSubmit             EventType = "submit"
	EventTypeSuspend            EventType = "suspend"
	EventTypeTimeUpdate         EventType = "timeupdate"
	EventTypeToggle             EventType = "toggle"
	EventTypeTouchCancel        EventType = "touchcancel"
	EventTypeTouchEnd           EventType = "touchend"
	EventTypeTouchMove          EventType = "touchmove"
	EventTypeTouchStart         EventType = "touchstart"
	EventTypeTransitionEnd      EventType = "transitionend"
	EventTypeUnload             EventType = "unload"
	EventTypeVolumeChange       EventType = "volumechange"
	EventTypeWaiting            EventType = "waiting"
	EventTypeWheel              EventType = "wheel"
)