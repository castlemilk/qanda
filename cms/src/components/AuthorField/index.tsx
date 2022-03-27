import React from "react";
import { TextField } from "@mui/material";
import { FieldDescription, FieldProps } from "@camberi/firecms";

interface AuthorFieldProps {
  author: string;
}

export default function CustomColorTextField({
  property,
  value,
  setValue,
  customProps,
  touched,
  error,
  isSubmitting,
  context, // the rest of the entity values here
  ...props
}: FieldProps<string, AuthorFieldProps>) {
  return (
    <>
      <TextField
        required={property.validation?.required}
        error={!!error}
        disabled={isSubmitting}
        label={property.title}
        value={value ?? context}
        onChange={(evt: any) => {
          setValue(evt.target.value);
        }}
        helperText={error}
        fullWidth
        variant={"filled"}
      />

      <FieldDescription property={property} />
    </>
  );
}
