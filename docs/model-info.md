## Command

```
saved_model_cli show --all --dir pkg/guesser/model
```

## Output

```
MetaGraphDef with tag-set: 'serve' contains the following SignatureDefs:

signature_def['classification']:
  The given SavedModel SignatureDef contains the following input(s):
    inputs['inputs'] tensor_info:
        dtype: DT_STRING
        shape: (-1)
        name: Placeholder:0
  The given SavedModel SignatureDef contains the following output(s):
    outputs['classes'] tensor_info:
        dtype: DT_STRING
        shape: (-1, 54)
        name: head/Tile:0
    outputs['scores'] tensor_info:
        dtype: DT_FLOAT
        shape: (-1, 54)
        name: head/predictions/probabilities:0
  Method name is: tensorflow/serving/classify

signature_def['predict']:
  The given SavedModel SignatureDef contains the following input(s):
    inputs['content'] tensor_info:
        dtype: DT_STRING
        shape: (-1)
        name: Placeholder:0
  The given SavedModel SignatureDef contains the following output(s):
    outputs['all_class_ids'] tensor_info:
        dtype: DT_INT32
        shape: (-1, 54)
        name: head/predictions/Tile:0
    outputs['all_classes'] tensor_info:
        dtype: DT_STRING
        shape: (-1, 54)
        name: head/predictions/Tile_1:0
    outputs['class_ids'] tensor_info:
        dtype: DT_INT64
        shape: (-1, 1)
        name: head/predictions/ExpandDims:0
    outputs['classes'] tensor_info:
        dtype: DT_STRING
        shape: (-1, 1)
        name: head/predictions/hash_table_Lookup/LookupTableFindV2:0
    outputs['logits'] tensor_info:
        dtype: DT_FLOAT
        shape: (-1, 54)
        name: add:0
    outputs['probabilities'] tensor_info:
        dtype: DT_FLOAT
        shape: (-1, 54)
        name: head/predictions/probabilities:0
  Method name is: tensorflow/serving/predict

signature_def['serving_default']:
  The given SavedModel SignatureDef contains the following input(s):
    inputs['inputs'] tensor_info:
        dtype: DT_STRING
        shape: (-1)
        name: Placeholder:0
  The given SavedModel SignatureDef contains the following output(s):
    outputs['classes'] tensor_info:
        dtype: DT_STRING
        shape: (-1, 54)
        name: head/Tile:0
    outputs['scores'] tensor_info:
        dtype: DT_FLOAT
        shape: (-1, 54)
        name: head/predictions/probabilities:0
  Method name is: tensorflow/serving/classify
The MetaGraph with tag set ['serve'] contains the following ops: {'Size', 'BiasAdd', 'ResourceGather', 'GatherV2', 'Placeholder', 'ArgMax', 'StringToHashBucketFast', 'AssignVariableOp', 'Add', 'NotEqual', 'SparseReshape', 'Tile', 'StridedSlice', 'StaticRegexFullMatch', 'If', 'ScalarSummary', 'SparseFillEmptyRows', 'StringJoin', 'SaveV2', 'LookupTableImportV2', 'Fill', 'LookupTableFindV2', 'ReadVariableOp', 'TensorListStack', 'Mul', 'TensorListFromTensor', 'Reshape', 'RandomUniform', 'PlaceholderWithDefault', 'ZerosLike', 'RestoreV2', 'Const', 'GatherNd', 'Identity', 'StatelessIf', 'TruncatedNormal', 'Prod', 'Unique', 'LessEqual', 'HistogramSummary', 'MergeV2Checkpoints', 'ExpandDims', 'VarHandleOp', 'TensorListReserve', 'GreaterEqual', 'Cast', 'MatMul', 'SparseSegmentMean', 'SparseSegmentSum', 'Select', 'NoOp', 'Range', 'StatelessWhile', 'Pack', 'Relu', 'VarIsInitializedOp', 'Sub', 'Shape', 'HashTableV2', 'AddV2', 'RealDiv', 'Where', 'Slice', 'ShardedFilename', 'Equal', 'Softmax', 'ConcatV2'}
```
