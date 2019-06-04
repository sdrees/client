import React from 'react'
import ClickableBox from './clickable-box'
import Text from './text'
import * as Styles from '../styles'
import {Props} from './radio-button'

export const RADIOBUTTON_SIZE = 22
export const RADIOBUTTON_MARGIN = 8

const RadioOuterCircle = Styles.styled(ClickableBox)(
  {
    backgroundColor: Styles.globalColors.white,
    borderRadius: 100,
    borderWidth: 1,
    height: RADIOBUTTON_SIZE,
    marginRight: RADIOBUTTON_MARGIN,
    position: 'relative' as 'relative',
    width: RADIOBUTTON_SIZE,
  },
  ({disabled, selected}) => ({
    borderColor: selected ? Styles.globalColors.blue : Styles.globalColors.black_20,
    opacity: disabled ? 0.4 : 1,
  })
)

const RadioInnerCircle = Styles.styled(ClickableBox)(
  {
    borderColor: Styles.globalColors.white,
    borderRadius: 10,
    borderWidth: 5,
    left: 5,
    position: 'absolute',
    top: 5,
  },
  ({selected}) => ({
    borderColor: selected ? Styles.globalColors.blue : Styles.globalColors.white,
  })
)

const RadioButton = ({disabled, label, onSelect, selected, style}: Props) => (
  <ClickableBox
    style={{...styleContainer, ...style}}
    onClick={disabled ? undefined : () => onSelect(!selected)}
  >
    {/*
    // @ts-ignore styled is more complicated than its typing lets on */}
    <RadioOuterCircle disabled={disabled} selected={selected}>
      {/*
      // @ts-ignore styled is more complicated than its typing lets on */}
      <RadioInnerCircle selected={selected} />
    </RadioOuterCircle>
    <Text type="Body" style={{color: Styles.globalColors.black}}>
      {label}
    </Text>
  </ClickableBox>
)

const styleContainer = {
  ...Styles.globalStyles.flexBoxRow,
  alignItems: 'center' as 'center',
  paddingBottom: Styles.globalMargins.xtiny,
  paddingTop: Styles.globalMargins.xtiny,
}

export default RadioButton
