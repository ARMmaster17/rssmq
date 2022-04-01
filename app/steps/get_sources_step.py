from watergrid.context import DataContext, OutputMode
from watergrid.steps import Step

from app.dto.source_dto import SourceDTO


class GetSourcesStep(Step):
    def __init__(self, sources: list):
        super().__init__(self.__class__.__name__, provides=['source'])
        self._sources = sources

    def run(self, context: DataContext):
        source_containers = []
        for source in self._sources:
            source_containers.append(SourceDTO(source))
        context.set('source', source_containers)
        context.set_output_mode(OutputMode.SPLIT)
