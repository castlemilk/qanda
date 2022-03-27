import React from "react";

import { Auth, User as FirebaseUser } from "firebase/auth";
import {
  Authenticator,
  buildCollection,
  buildProperty,
  buildSchema,
  buildEntityCallbacks,
  EntityReference,
  FirebaseCMSApp,
  EntityCallbacks,
  NavigationBuilder,
  EntityOnSaveProps,
  NavigationBuilderProps,
} from "@camberi/firecms";

import "typeface-rubik";
import "typeface-space-mono";

import { Question } from "./lib/question/v1alpha1/question_pb";
import { Author } from "./lib/api/v1alpha1/question_pb";

const firebaseConfig = {
  apiKey: "AIzaSyDiLpbL8wrFEjsNxVuNCX2D03Mx5eQNK2s",
  authDomain: "high-keel-345423.firebaseapp.com",
  projectId: "high-keel-345423",
  storageBucket: "high-keel-345423.appspot.com",
  messagingSenderId: "849980621434",
  appId: "1:849980621434:web:fc812924930165890231c3",
  measurementId: "G-B09VJZRR0K",
};

const locales = {
  "en-US": "English (United States)",
  "es-ES": "Spanish (Spain)",
  "de-DE": "German",
};

const productSchema = buildSchema<Question.AsObject>({
  name: "Question",
  properties: {
    metadata: {
      title: "Metadata",
      validation: { required: true },
      dataType: "map",
      properties: {
        name: {
          title: "Name",
          dataType: "string",
        },
        createdat: {
          title: "Created at",
          dataType: "timestamp",
          autoValue: "on_create",
        },
      },
    },
    author: {
      title: "Author",
      dataType: "string",
      disabled: true,
    },
    spec: {
      title: "Spec",
      validation: { required: true },
      dataType: "map",
      properties: {
        content: {
          title: "Question",
          dataType: "string",
          description: "The contents of your Question",
          longDescription:
            "Here you can populate a question which will then be viewable in the QandA website. You'll be able to provide a corresponding answer in the subsequent fields",
        },
        tagsList: {
          title: "tags",
          description: "provide tags to make your question more discoverable",
          dataType: "array",
          of: {
            dataType: "string",
          },
        },
      },
    },
  },
});

const authorCallback = buildEntityCallbacks<Question.AsObject>({
  onPreSave: ({ values, context }) => {
    values.author = context.authController.user.email;
    return values;
  },
  onSaveSuccess: (props: EntityOnSaveProps<Question.AsObject>) => {
    console.log("onSaveSuccess", props);
  },

  onSaveFailure: (props: EntityOnSaveProps<Question.AsObject>) => {
    console.log("onSaveFailure", props);
  },
});

// const localeSchema = buildSchema({
//     customId: locales,
//     name: "Locale",
//     properties: {
//         title: {
//             title: "Title",
//             validation: { required: true },
//             dataType: "string"
//         },
//         selectable: {
//             title: "Selectable",
//             description: "Is this locale selectable",
//             dataType: "boolean"
//         }
//     }
// });

export default function App() {
  const navigation: NavigationBuilder = async ({
    user,
    authController,
  }: NavigationBuilderProps) => {
    return {
      collections: [
        buildCollection({
          path: "qanda",
          schema: productSchema,
          name: "QandA",
          permissions: ({ authController }) => ({
            edit: true,
            create: true,
            // we have created the roles object in the navigation builder
            delete: authController.extra.roles.includes("admin"),
          }),
          callbacks: authorCallback,
        }),
      ],
    };
  };

  const myAuthenticator: Authenticator<FirebaseUser> = async ({
    user,
    authController,
  }) => {
    // You can throw an error to display a message
    if (user?.email?.includes("flanders")) {
      throw Error("Stupid Flanders!");
    }

    console.log("Allowing access to", user?.email);
    // This is an example of retrieving async data related to the user
    // and storing it in the user extra field.
    const sampleUserData = await Promise.resolve({
      roles: ["admin"],
    });
    authController.setExtra(sampleUserData);
    return true;
  };

  return (
    <FirebaseCMSApp
      name={"QandA creation"}
      authentication={myAuthenticator}
      navigation={navigation}
      firebaseConfig={firebaseConfig}
    />
  );
}
